package mongodb

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key struct {
	streamName string
	shardID    string
}

// Option is used to override defaults when creating a new Checkpoint
type Option func(*Checkpoint)

// WithMaxInterval sets the flush interval
func WithMaxInterval(maxInterval time.Duration) Option {
	return func(c *Checkpoint) {
		c.maxInterval = maxInterval
	}
}

// Checkpoint stores and retrieves the last evaluated key
type Checkpoint struct {
	appName     string
	conn        *mongo.Collection
	mu          *sync.Mutex // protects the checkpoints
	done        chan struct{}
	checkpoints map[key]string
	maxInterval time.Duration
}

// New returns a checkpoint that uses mongoDB for underlying storage
// Using connectionStr turn it more flexible to use specific db configs
func New(appName string, conn *mongo.Collection, opts ...Option) (*Checkpoint, error) {
	if appName == "" {
		return nil, errors.New("application name not defined")
	}

	ck := &Checkpoint{
		appName:     appName,
		conn:        conn,
		done:        make(chan struct{}),
		maxInterval: 1 * time.Minute,
		mu:          new(sync.Mutex),
		checkpoints: map[key]string{},
	}

	for _, opt := range opts {
		opt(ck)
	}

	go ck.loop()

	return ck, nil
}

// GetMaxInterval returns the maximum interval before the checkpoint
func (c *Checkpoint) GetMaxInterval() time.Duration {
	return c.maxInterval
}

type appCheckpoint struct {
	Namespace      string `bson:"namespace"`
	ShardID        string `bson:"shard_id"`
	SequenceNumber string
}

// GetCheckpoint determines if a checkpoint for a particular Shard exists.
// Typically used to determine whether we should start processing the shard with
// TRIM_HORIZON or AFTER_SEQUENCE_NUMBER (if checkpoint exists).
func (c *Checkpoint) GetCheckpoint(streamName, shardID string) (string, error) {
	namespace := fmt.Sprintf("%s-%s", c.appName, streamName)
	filter := bson.D{{"namespace", namespace}, {"shard_id", shardID}}

	var result appCheckpoint
	err := c.conn.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		logger.Error("kinesis mongodb sotre - GetCheckpoint: ", err.Error())
		return "", err
	}

	return result.SequenceNumber, nil
}

// SetCheckpoint stores a checkpoint for a shard (e.g. sequence number of last record processed by application).
// Upon failover, record processing is resumed from this point.
func (c *Checkpoint) SetCheckpoint(streamName, shardID, sequenceNumber string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if sequenceNumber == "" {
		return fmt.Errorf("sequence number should not be empty")
	}

	key := key{
		streamName: streamName,
		shardID:    shardID,
	}

	c.checkpoints[key] = sequenceNumber

	return nil
}

// Shutdown the checkpoint. Save any in-flight data.
func (c *Checkpoint) Shutdown() error {
	c.done <- struct{}{}

	return c.save()
}

func (c *Checkpoint) loop() {
	tick := time.NewTicker(c.maxInterval)
	defer tick.Stop()
	defer close(c.done)

	for {
		select {
		case <-tick.C:
			c.save()
		case <-c.done:
			return
		}
	}
}

func (c *Checkpoint) save() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, sequenceNumber := range c.checkpoints {
		namespace := fmt.Sprintf("%s-%s", c.appName, key.streamName)
		filter := bson.D{{"namespace", namespace}, {"shard_id", key.shardID}}

		update := bson.M{
			"$set": bson.M{
				"namespace":      namespace,
				"shardID":        key.shardID,
				"sequenceNumber": sequenceNumber,
			},
		}

		updateResult, err := c.conn.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
		if err != nil {
			logger.Error("kinesis mongodb sotre - save: ", err.Error())
			return err
		}
		logger.Info("kinesis mongodb sotre - save: ", updateResult)
	}

	return nil
}
