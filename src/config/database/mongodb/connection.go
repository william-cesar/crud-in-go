package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	conn, db := os.Getenv("DB_CONNECTION"), os.Getenv("DB_NAME")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return client.Database(db), nil
}
