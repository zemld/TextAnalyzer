package handlers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName        string = "mg"
	connectionURI string = "mongodb://root:password@localhost:27017"
)

const (
	filesCollection     string = "files"
	wordCloudCollection string = "wordcloud"
)

func storeDocument(f []byte, id int, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
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

func getDocument(id int, collectionName string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	var result bson.M
	err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["file"].([]byte), nil
}
