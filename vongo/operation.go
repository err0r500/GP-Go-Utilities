package vongo

import (
	"context"

	"github.com/VoodooTeam/GP-Go-Utilities/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create method
func Create(ctx context.Context, d Document, opts ...*options.InsertOneOptions) error {
	collName := GetCollectionName(d)

	// Call to saving hook
	if err := callToBeforeSaveHook(d, true); err != nil {
		return err
	}

	if ccw, ok := DBConn.Config.(ConnectionContextWrapperHook); ok {
		ctx = ccw.WrapContext(ctx)
	}

	coll := DBConn.Collection(collName)
	insertOneResult, err := coll.InsertOne(ctx, d)
	if err != nil {
		return err
	}
	if insertOneResult != nil {
		if id, ok := insertOneResult.InsertedID.(primitive.ObjectID); ok {
			d.SetID(id)
		} else {
			logger.Errorf("Create - Could not parse insertOneResult.InsertedID: %v", insertOneResult.InsertedID)
		}
	} else {
		logger.Errorf("Create - insertOneResult is nil")
	}

	return callToAfterSaveHook(d)
}

// UpdateOne method
func UpdateOne(ctx context.Context, d Document, filter interface{}, opts ...*options.UpdateOptions) error {
	collName := GetCollectionName(d)

	isNew := true
	if newt, ok := d.(NewTrackerHook); ok {
		isNew = newt.IsNew()
	}

	// Call to saving hook
	if err := callToBeforeSaveHook(d, isNew); err != nil {
		return err
	}

	if ccw, ok := DBConn.Config.(ConnectionContextWrapperHook); ok {
		ctx = ccw.WrapContext(ctx)
	}

	coll := DBConn.Collection(collName)
	updateResult, err := coll.UpdateOne(ctx, filter, bson.M{"$set": d}, opts...)
	if err != nil {
		return err
	}
	if updateResult != nil {
		if id, ok := updateResult.UpsertedID.(primitive.ObjectID); ok {
			d.SetID(id)
		}
	}

	return callToAfterSaveHook(d)
}

// UpdateFields method
func UpdateFields(ctx context.Context, d Document, filter interface{}, update func(d Document) interface{}, opts ...*options.UpdateOptions) error {
	collName := GetCollectionName(d)

	isNew := true
	if newt, ok := d.(NewTrackerHook); ok {
		isNew = newt.IsNew()
	}

	// Call to saving hook
	if err := callToBeforeSaveHook(d, isNew); err != nil {
		return err
	}

	if ccw, ok := DBConn.Config.(ConnectionContextWrapperHook); ok {
		ctx = ccw.WrapContext(ctx)
	}

	updateBson := update(d)
	switch ubson := updateBson.(type) {
	case bson.D:
		ubson = append(ubson, bson.E{Key: "$currentDate", Value: bson.D{{Key: "updated_at", Value: true}}})
		updateBson = ubson
	case bson.M:
		if currentDate, ok := ubson["$currentDate"]; ok {
			currentDate.(bson.M)["updated_at"] = true
			ubson["$currentDate"] = currentDate
		} else {
			ubson["$currentDate"] = bson.M{
				"updated_at": true,
			}
		}
		updateBson = ubson
	}

	coll := DBConn.Collection(collName)
	updateResult, err := coll.UpdateOne(ctx, filter, updateBson, opts...)
	if err != nil {
		return err
	}
	if updateResult != nil {
		if id, ok := updateResult.UpsertedID.(primitive.ObjectID); ok {
			d.SetID(id)
		}
	}

	return callToAfterSaveHook(d)
}

// FindByID method
func FindByID(ctx context.Context, d Document, opts ...*options.FindOneOptions) error {
	return First(ctx, d.BsonID(), d, opts...)
}

// First method
func First(ctx context.Context, filter interface{}, d Document, opts ...*options.FindOneOptions) error {
	collName := GetCollectionName(d)

	if ccw, ok := DBConn.Config.(ConnectionContextWrapperHook); ok {
		ctx = ccw.WrapContext(ctx)
	}

	coll := DBConn.Collection(collName)
	err := coll.FindOne(ctx, filter, opts...).Decode(d)
	if err != nil {
		return err
	}

	return callToAfterFindHook(d)
}

// Find method
func Find(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	var collName string
	var err error

	if hook, ok := results.(GetCollectionNameHook); ok {
		collName = hook.GetCollectionName()
	} else {
		collName, err = getCollectionName(results)
		if err != nil {
			return err
		}
	}

	if ccw, ok := DBConn.Config.(ConnectionContextWrapperHook); ok {
		ctx = ccw.WrapContext(ctx)
	}

	coll := DBConn.Collection(collName)
	cur, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}
