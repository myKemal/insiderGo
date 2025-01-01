package initialize

import (
	"context"
	"fmt"
	"github.com/myKemal/insiderGo/app/config"
	"github.com/myKemal/insiderGo/app/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Mongo() (repository.MongoRepository, error) {

	client, err := connectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	return repository.NewMongoRepository(client, "insider_chat", "message"), nil
}

func connectToMongoDB() (*mongo.Client, error) {
	uri := config.EnvMongoURI()
	if uri == "" {
		log.Fatal("MONGODB_URI not found in environment")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return client, nil
}
