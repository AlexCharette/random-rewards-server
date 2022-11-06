package api

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Rewards *mongo.Collection
)

func Run() {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("No environment variable found named DB_URI")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	Rewards = client.Database(os.Getenv("DB_NAME")).Collection("rewards")
}
