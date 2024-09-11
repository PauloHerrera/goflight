package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "user_flights"

func (db *Worker) GetFlight() (results []bson.M, err error) {
	coll := db.client.Database(db.database).Collection(collectionName)

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
		return
	}

	return results, err
}

func (db *Worker) PostUserFlights(ctx context.Context, userFlights *UserFlight) (results mongo.InsertOneResult, err error) {
	coll := db.client.Database(db.database).Collection(collectionName)

	_, err = coll.InsertOne(ctx, userFlights)
	if err != nil {
		log.Fatal("failed to insert flight data", err)
		return
	}

	return
}
