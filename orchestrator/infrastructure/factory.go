package factory

import (
	"fmt"

	mongoFactory "github.com/bharath0292/quantdrey/infrastructure/mongo"
	natsfactory "github.com/bharath0292/quantdrey/infrastructure/nats"
	postgresFactory "github.com/bharath0292/quantdrey/infrastructure/postgres"
	redisFactory "github.com/bharath0292/quantdrey/infrastructure/redis"
	"github.com/rs/zerolog/log"
)

type FactoryConfig struct {
	RedisConfig    redisFactory.RedisConfig
	NatsConfig     natsfactory.NatsConfig
	PostgresConfig postgresFactory.PostgresConfig
	MongoConfig    mongoFactory.MongoConfig
}

type Factory struct {
	RedisClient    *redisFactory.RedisClient
	NatsClient     *natsfactory.NatsClient
	PostgresClient *postgresFactory.PostgresClient
	MongoClient    *mongoFactory.MongoClient
}

func NewFactory(configs FactoryConfig) (*Factory, error) {

	/* redis connection */
	redisClient, err := redisFactory.NewRedisClient(configs.RedisConfig)
	if err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}
	log.Info().Msg("Redis connection established")

	/* nats connection */
	natsClient, err := natsfactory.NewNatsClient(configs.NatsConfig)
	if err != nil {
		return nil, fmt.Errorf("nats connection failed: %w", err)
	}
	log.Info().Msg("Nats connection established")

	/* postgres connection */
	postgresClient, err := postgresFactory.NewPostgresClient(configs.PostgresConfig)
	if err != nil {
		return nil, fmt.Errorf("postgres connection failed: %w", err)
	}
	log.Info().Msg("Postgres connection established")

	/* postgres connection */
	mongoClient, err := mongoFactory.NewMongoClient(configs.MongoConfig)
	if err != nil {
		return nil, fmt.Errorf("mongo connection failed: %w", err)
	}
	log.Info().Msg("Mongo connection established")

	return &Factory{
		RedisClient:    redisClient,
		NatsClient:     natsClient,
		PostgresClient: postgresClient,
		MongoClient:    mongoClient,
	}, nil
}

func (f *Factory) Close() error {
	f.RedisClient.Close()
	f.NatsClient.Close()
	f.PostgresClient.Close()

	if err := f.MongoClient.Close(); err != nil {
		return fmt.Errorf("mongo close failed: %w", err)
	}

	log.Info().Msg("Factories disconnected")
	return nil
}
