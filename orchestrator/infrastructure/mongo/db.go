package mongoFactory

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Uri         string
	Username    *string
	Password    *string
	MaxPoolSize *uint64
	MinPoolSize *uint64
}

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(ctx context.Context, config MongoConfig) (*MongoClient, error) {
	timeout := time.Duration(3) * time.Second

	var sb strings.Builder
	sb.WriteString("mongodb://")
	sb.WriteString(*config.Username)
	sb.WriteString(":")
	sb.WriteString(*config.Password)
	sb.WriteString("@")
	sb.WriteString(config.Uri)
	sb.WriteString("/")

	fullMongoUri := sb.String()

	clientOptions := options.Client().ApplyURI(fullMongoUri)
	clientOptions.MaxPoolSize = config.MaxPoolSize
	clientOptions.MinPoolSize = config.MinPoolSize
	clientOptions.Timeout = &timeout

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &MongoClient{client}, nil
}
