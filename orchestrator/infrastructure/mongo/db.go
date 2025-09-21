package mongoFactory

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Uri         string
	Username    string
	Password    string
	Database    string
	MaxPoolSize uint64
	MinPoolSize uint64
}

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoClient(config *MongoConfig) (*MongoClient, error) {
	if config == nil {
		return nil, errors.New("missing config")
	}

	timeout := time.Duration(30) * time.Second

	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s/", config.Username, config.Password, config.Uri)

	clientOptions := options.Client().ApplyURI(mongoUri)
	clientOptions.Timeout = &timeout
	clientOptions.MaxPoolSize = &config.MaxPoolSize
	clientOptions.MinPoolSize = &config.MinPoolSize

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &MongoClient{client, client.Database(config.Database)}, nil
}

func (c *MongoClient) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("close connection failed: %w", err)
	}
	return nil
}

func (c *MongoClient) ReadDocument(ctx context.Context, collection string, id bson.ObjectID, out any) error {
	if out == nil {
		return errors.New("output object not available")
	}

	coll := c.db.Collection(collection)
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(out)
	if err != nil {
		log.Error().Err(err).
			Str("collection", collection).
			Str("insertedId", id.Hex()).
			Msg("Failed to fetch inserted document")
		return mongo.ErrNilDocument
	}

	return nil
}

func (c *MongoClient) CreateDocument(ctx context.Context, collection string, document any, out any) (bson.ObjectID, error) {
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

func (c *MongoClient) UpdateDocument(ctx context.Context, collection string, filter bson.M, update bson.M, out any) (bson.ObjectID, error) {
	coll := c.db.Collection(collection)

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error().Err(err).Str("collection", collection).Msg("Failed to update document")
		return bson.NilObjectID, err
	}

	if result.MatchedCount == 0 {
		log.Error().Str("collection", collection).Msg("No matching document found for update")
		return bson.NilObjectID, mongo.ErrNoDocuments
	}

	oid := filter["_id"].(bson.ObjectID)

	// Optionally fetch updated document
	if out != nil {
		err = coll.FindOne(ctx, bson.M{"_id": oid}).Decode(out)
		if err != nil {
			log.Error().Err(err).
				Str("collection", collection).
				Str("updatedId", oid.Hex()).
				Msg("Failed to fetch updated document")
			return bson.NilObjectID, err
		}
	}

	return oid, nil
}
