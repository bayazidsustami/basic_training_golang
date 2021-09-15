package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func main() {

}

func connect() (*mongo.Database, error) {
	clientOption := options.Client()
	clientOption.ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}
