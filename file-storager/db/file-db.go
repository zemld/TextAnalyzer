package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName string = "mg"
)

const (
	FilesCollection     string = "files"
	WordCloudCollection string = "wordcloud"
)

func StoreDocument(f []byte, id int, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	_, err = db.Collection(collectionName).InsertOne(ctx, bson.M{"_id": id, "file": f})
	if err != nil {
		return err
	}
	return nil
}
