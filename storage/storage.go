package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Worker struct {
	database string
	client   *mongo.Client
}

func NewWorker(databaseUri string, database string) *Worker {
	client, err := mongo.Connect(context.TODO(), options.Client(), options.Client().ApplyURI(databaseUri))

	if err != nil {
		log.Fatal("failed to connect connect database:", err)
	}

	return &Worker{
		client:   client,
		database: database,
	}
}
