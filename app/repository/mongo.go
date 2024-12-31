package repository

import (
	"context"
	"github.com/myKemal/mongoApi/app/daos"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
	FetchUnsentMessages() ([]daos.MessageDAO, error)
}
type mongoRepositoryImpl struct {
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, dbName, collectionName string) MongoRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &mongoRepositoryImpl{collection: collection}
}

func (r *mongoRepositoryImpl) FetchUnsentMessages() ([]daos.MessageDAO, error) {
	filter := bson.M{"sending_status": "not_sent"}
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Error finding unsent messages: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var messages []daos.MessageDAO
	if err := cursor.All(context.TODO(), &messages); err != nil {
		log.Printf("Error decoding messages: %v", err)
		return nil, err
	}
	return messages, nil
}
