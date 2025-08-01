package main

import (
	"context"
	"os"
	"time"

	// Factories
	factory "github.com/bharath0292/quantdrey/infrastructure"
	mongoFactory "github.com/bharath0292/quantdrey/infrastructure/mongo"
	natFactory "github.com/bharath0292/quantdrey/infrastructure/nats"
	postgresFactory "github.com/bharath0292/quantdrey/infrastructure/postgres"
	redisFactory "github.com/bharath0292/quantdrey/infrastructure/redis"

	strategyservice "github.com/bharath0292/quantdrey/internal/domains/strategy/service"
	tickhub "github.com/bharath0292/quantdrey/internal/domains/tick/hub"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"

	engine "github.com/bharath0292/quantdrey/app/engines"

	util "github.com/bharath0292/quantdrey/pkg/utils"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// To set global timezone
	os.Setenv("TZ", "UTC")

	err := godotenv.Load()
	if err != nil {
		log.Info().Msgf("Error loading .env file: %v", err)
	}
	requiredEnvVars := []string{
		"GIN_MODE", "API_AUTH_KEY",
		"MONGO_URI", "MONGO_USERNAME", "MONGO_PASSWORD",
		"REDIS_URI", "REDIS_USERNAME", "REDIS_PASSWORD",
		"NATS_URI", "NATS_USERNAME", "NATS_PASSWORD",
	}
	if err := util.ValidateEnvVars(requiredEnvVars); err != nil {
		log.Logger.Fatal().Msg(err.Error())
		os.Exit(1)
	}

	logLevel := zerolog.InfoLevel
	if os.Getenv("GIN_MODE") == "debug" {
		logLevel = zerolog.DebugLevel
	}
	util.InitLogger(logLevel)
}

func main() {
	log.Info().Msg("Starting server")

	const (
		certPath = "certs/orchestrator.local.pem"
		keyPath  = "certs/orchestrator.local-key.pem"
		addr     = "orchestrator.local:8443"
	)

	if err := util.IsSelfSignedCertAvailable(certPath, keyPath); err != nil {
		log.Fatal().Err(err).Msg("Please run 'make create-certs' to generate a self-signed certificate")
	}

	redisUri := os.Getenv("REDIS_URI")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisOptions := redisFactory.RedisConfig{
		InitAddress:         []string{redisUri},
		Username:            redisUsername,
		Password:            redisPassword,
		BlockingPoolCleanup: time.Duration(5 * time.Minute),
		BlockingPoolMinSize: 10,
		BlockingPoolSize:    100,
	}

	natsUri := os.Getenv("NATS_URI")
	natsUsername := os.Getenv("NATS_USERNAME")
	natsPassword := os.Getenv("NATS_PASSWORD")
	natsOptions := natFactory.NatsConfig{
		URL:      natsUri,
		PoolSize: 10,
		Options: []nats.Option{
			nats.UserInfo(natsUsername, natsPassword),
		},
	}

	mongoUri := os.Getenv("MONGO_URI")
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	postgresConfig := postgresFactory.PostgresConfig{}
	mongoConfig := mongoFactory.MongoConfig{
		Uri:      mongoUri,
		Username: &mongoUsername,
		Password: &mongoPassword,
	}

	ctx := context.Background()
	factories, err := factory.NewFactory(ctx, factory.FactoryConfig{
		RedisConfig:    redisOptions,
		NatsConfig:     natsOptions,
		PostgresConfig: postgresConfig,
		MongoConfig:    mongoConfig,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Factory initialization failed")
	}

	/* ######### HUBS ######### */
	tickHub := tickhub.NewTickHub(factories.RedisClient, factories.NatsClient)
	tickHub.SubscribeTicks()

	/* ######### SERVICES ######### */
	tickService := tickservice.NewTickService(factories.NatsClient, tickHub)
	strategyService := strategyservice.NewStrategyService(tickService)

	/* ######### HANDLERS ######### */

	/* ######### Engines ######### */
	restEngine := engine.NewRestEngine()
	graphqlEngine := engine.NewGraphQLEngine(strategyService)
	handlers := engine.RestEngineHandlers{}
	restEngine.Setup(handlers, graphqlEngine)

	log.Info().Msgf("Starting HTTPS server at https://%s", addr)
	if err := restEngine.RunTLS(":8443", certPath, keyPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTPS server")
	}
}
