package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (db *Worker) PutUserFlights(ctx context.Context, userFlights *UserFlight) (err error) {
	coll := db.client.Database(db.database).Collection(collectionName)

	filter := bson.M{"user_id": userFlights.UserID}
	update := bson.M{"$set": userFlights}

	_, err = coll.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	if err != nil {
		log.Fatal("failed to insert flight data", err)
		return
	}

	return
}
