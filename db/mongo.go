package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	clientOpt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	col := client.Database("cms").Collection("customer")
	return col
}
