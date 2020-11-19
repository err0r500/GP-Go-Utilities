package vongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection struct
type Connection struct {
	Config   ConfigInterface
	client   *mongo.Client
	database *mongo.Database
}

// DBConn is the connection initialized after Connect is called.
// All underlying operations are made using this connection
var DBConn *Connection

// InitConnection method
func InitConnection(config ConfigInterface) (*Connection, error) {
	conn := &Connection{
		Config: config,
	}

	err := conn.Connect()

	if err != nil {
		DBConn = nil
		log.Printf("Error while connectiong to MongoDb (err: %v)", err)

		return nil, err
	}

	conn.database = conn.client.Database(config.GetDatabase())
	DBConn = conn
	return conn, err
}

// Connect to the database using the provided config
func (conn *Connection) Connect() (err error) {
	defer func() {
		if r := recover(); r != nil {
			// panic(r)
			// return
			if e, ok := r.(error); ok {
				err = e
			} else if e, ok := r.(string); ok {
				err = errors.New(e)
			} else {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()

	clientOptions := options.Client().ApplyURI(conn.Config.GetURI())
	if conn.Config.GetMonitor() != nil {
		clientOptions = clientOptions.SetMonitor(conn.Config.GetMonitor())
	}

	conn.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = conn.client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	logger.Debug("Connected to MongoDB!")
	return nil
}

// Collection method
func (conn *Connection) Collection(name string) *mongo.Collection {
	return conn.database.Collection(name)
}
