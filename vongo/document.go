package vongo

import (
	"context"
	"errors"
	"time"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetCollectionNameHook get collection name
type GetCollectionNameHook interface {
	GetCollectionName() string
}

// Document interface
type Document interface {
	TimeCreatedAtTrackerHook
	TimeUpdatedAtTrackerHook

	GetID() interface{}
	SetID(interface{})

	BsonID() *bson.M
}

// DocumentModel struct
type DocumentModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	// Model lifecycle flags
	exists bool `bson:"-"`
}

// GetCreatedAt gets the created date
func (d *DocumentModel) GetCreatedAt() time.Time {
	return d.CreatedAt
}

// SetCreatedAt sets the created date
func (d *DocumentModel) SetCreatedAt(t time.Time) {
	d.CreatedAt = t
}

// GetUpdatedAt gets the modified date
func (d *DocumentModel) GetUpdatedAt() time.Time {
	return d.UpdatedAt
}

// SetUpdatedAt sets the modified date
func (d *DocumentModel) SetUpdatedAt(t time.Time) {
	d.UpdatedAt = t
}

// IsNew to ask Is the document new
func (d *DocumentModel) IsNew() bool {
	return d.ID.IsZero()
}

/**
** The Document interface implementation
**/

// GetID satisfies the document interface
func (d *DocumentModel) GetID() interface{} {
	if d.ID.IsZero() {
		return nil
	}

	return d.ID
}

// SetID sets the ID for the document
func (d *DocumentModel) SetID(id interface{}) {
	if id, ok := id.(primitive.ObjectID); ok {
		d.ID = id
	} else {
		logger.Errorf("DocumentModel cannot set ID")
	}
}

// BsonID returns the document id using bson.M interface style
// This method can be directly used with Find, but not with FindID
// which expects directly id interface{} (i.e. d.ID/d.GetID())
func (d *DocumentModel) BsonID() *bson.M {
	return &bson.M{
		"_id": d.GetID(),
	}
}

// Save method
func Save(ctx context.Context, d Document) error {
	// If the model implements the NewTracker interface, we'll use that to determine newness. Otherwise always assume it's new
	isNew := true
	if newt, ok := d.(NewTrackerHook); ok {
		isNew = newt.IsNew()
	}

	id := d.GetID()
	if !isNew && id != nil {
		return errors.New("New tracker says this document isn't new but there is no valid Id field")
	}

	if isNew {
		if err := Create(ctx, d); err != nil {
			return err
		}
	} else {
		if err := UpdateOne(ctx, d, d.BsonID()); err != nil {
			return err
		}
	}

	return nil
}

// Remove removes document from database, running
// before and after delete hooks
func Remove(ctx context.Context, d Document) error {
	// Create a new session per mgo's suggestion to avoid blocking
	collName := GetCollectionName(d)

	if err := callToBeforeRemoveHook(d); err != nil {
		return err
	}

	if ccw, ok := DBConn.Config.(ConnectionContextWrapperHook); ok {
		ctx = ccw.WrapContext(ctx)
	}

	coll := DBConn.Collection(collName)
	if _, err := coll.DeleteOne(ctx, d.BsonID()); err != nil {
		return err
	}

	if err := callToAfterRemoveHook(d); err != nil {
		return err
	}

	return nil
}

// GetCollectionName method
func GetCollectionName(d Document) string {
	if hook, ok := d.(GetCollectionNameHook); ok {
		return hook.GetCollectionName()
	}

	collName, err := getCollectionName(d)
	if err != nil {
		panic("Document don't have collection name")
	}
	return collName
}
