package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() {
	conn := os.Getenv("DB_CONNECTION")

	if conn == "" {
		log.Fatal("Database connection string is not set")
		return
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB")
}
