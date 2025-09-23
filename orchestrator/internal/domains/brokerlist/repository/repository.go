package brokerlistrepository

import (
	"context"
	"fmt"

	postgresFactory "github.com/bharath0292/quantdrey/infrastructure/postgres"
	redisFactory "github.com/bharath0292/quantdrey/infrastructure/redis"
	brokersentity "github.com/bharath0292/quantdrey/internal/domains/brokerlist/entity"
)

type brokersRepository struct {
	postgre *postgresFactory.PostgresClient
	cache   *redisFactory.RedisClient
}

type IBrokersRepository interface {
	GetBrokers(ctx context.Context) ([]*brokersentity.Broker, error)
}

func NewBrokersRepository(
	postgre *postgresFactory.PostgresClient,
	cache *redisFactory.RedisClient,
) IBrokersRepository {
	return &brokersRepository{
		postgre: postgre,
		cache:   cache,
	}
}

func (b *brokersRepository) GetBrokers(ctx context.Context) ([]*brokersentity.Broker, error) {

	var brokers []*brokersentity.Broker

	if err := b.postgre.Select(ctx, nil, &brokers); err != nil {
		return nil, fmt.Errorf("failed to get brokers list: %w", err)
	}
	return brokers, nil
}
