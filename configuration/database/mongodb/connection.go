package mongodb

import (
	"context"
	"fullcycle-auction_go/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDatabase := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Error trying to connect to mongodb database", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to ping mongodb database", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
