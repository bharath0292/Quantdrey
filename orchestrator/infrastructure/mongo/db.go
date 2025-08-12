package mongoFactory

import (
	"context"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	db     *mongo.Database
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
	return &MongoClient{client, client.Database("quantdrey")}, nil
}

func (c *MongoClient) CreateDocument(ctx context.Context, collection string, document any, out any) (bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := c.db.Collection(collection)

	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		log.Error().Err(err).Str("collection", collection).Msg("Failed to insert document")
		return bson.NilObjectID, err
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		log.Error().Msg("InsertedID is not an ObjectID")
		return bson.NilObjectID, mongo.ErrNilDocument
	}

	if out != nil {
		err = coll.FindOne(ctx, bson.M{"_id": oid}).Decode(out)
		if err != nil {
			log.Error().Err(err).
				Str("collection", collection).
				Str("insertedId", oid.Hex()).
				Msg("Failed to fetch inserted document")
			return bson.NilObjectID, err
		}
	}

	return oid, nil
}

func (c *MongoClient) UpdateDocument() error {
	return nil
}
