package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"khot-queue-api/configs"
	"log"
	"time"
)

var mongoClient *mongo.Client
var DB *mongo.Database

func Connect() (*mongo.Database, error) {
	jsonConfig := configs.LoadJsonConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(jsonConfig.DbConnection)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to MongoDB!")

	mongoClient = client
	DB = client.Database("khot-queue")
	// ระบุชื่อ database ที่จะใช้
	return client.Database("khot-queue"), nil
}

// DisconnectMongo ปิด connection อย่างปลอดภัย
func DisconnectMongo() error {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return mongoClient.Disconnect(ctx)
	}
	return nil
}
