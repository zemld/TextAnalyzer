package handlers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
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

func storeDocument(f []byte, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	_, err = db.Collection(filesCollection).InsertOne(ctx, bson.M{"_id": id, "file": f})
	if err != nil {
		return err
	}
	return nil
}

func getDocument(id int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	var result bson.M
	err = db.Collection(filesCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["file"].([]byte), nil
}

func storeWordCloud(id int, file multipart.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	db := client.Database(dbName)
	options := options.GridFSBucket().SetName(wordCloudCollection)
	bucket, err := gridfs.NewBucket(db, options)
	if err != nil {
		return err
	}

	uploadStream, err := bucket.OpenUploadStreamWithID(id, fmt.Sprintf("%d.png", id))
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	if _, err = io.Copy(uploadStream, file); err != nil {
		return err
	}

	return nil
}

func getWordCloud(id int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	options := options.GridFSBucket().SetName(wordCloudCollection)
	bucket, err := gridfs.NewBucket(db, options)
	if err != nil {
		return nil, err
	}
	downloadStream, err := bucket.OpenDownloadStream(id)
	if err != nil {
		return nil, err
	}
	defer downloadStream.Close()

	data, err := io.ReadAll(downloadStream)
	if err != nil {
		return nil, err
	}
	return data, nil
}
