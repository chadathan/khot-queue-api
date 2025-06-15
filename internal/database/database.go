package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"khot-queue-api/configs"
	"log"
)

func Connect() (*mongo.Client, error) {
	jsonConfig := configs.LoadJsonConfig()

	clientOptions := options.Client().ApplyURI(jsonConfig.DbConnection)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
