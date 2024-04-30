package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	generatorConnectionName = "generator"
)

type autoIncrementID struct {
	ID int64 `bson:"seq" json:"seq"`
}

func InstallNextID(db *mongo.Database, key string, id int64) error {
	counterColl := db.Collection(generatorConnectionName)

	options := options.
		FindOneAndUpdate().
		SetUpsert(true)

	var result autoIncrementID
	err := counterColl.
		FindOneAndUpdate(
			context.Background(),
			bson.M{"_id": key},
			bson.M{SET: bson.M{"seq": id}},
			options).
		Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

func GetNextID(db *mongo.Database, key string) (int64, error) {
	counterColl := db.Collection(generatorConnectionName)

	options := options.
		FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var result autoIncrementID
	err := counterColl.
		FindOneAndUpdate(
			context.Background(),
			bson.M{"_id": key},
			bson.M{INC: bson.M{"seq": 1}},
			options).
		Decode(&result)
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}
