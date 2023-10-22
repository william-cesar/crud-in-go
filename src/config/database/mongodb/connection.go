package mongodb

import (
	"context"
	"os"

	"github.com/william-cesar/crud-in-go/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	logger.NewInfoLog(logger.JOURNEY["DB"], logger.MESSAGE["INIT"]["DB"])

	conn, db := os.Getenv("DB_CONNECTION"), os.Getenv("DB_NAME")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))

	if err != nil {
		logger.NewErrorLog(logger.JOURNEY["DB"], logger.MESSAGE["ERROR"]["DB_CONN"])
		return nil, err
	}

	// if err := client.Ping(ctx, nil); err != nil {
	// 	logger.NewErrorLog(logger.JOURNEY["DB"], logger.MESSAGE["ERROR"]["DB_PING"])
	// 	return nil, err
	// }

	logger.NewInfoLog(logger.JOURNEY["DB"], logger.MESSAGE["OK"]["DB_CONN"])

	return client.Database(db), nil
}
