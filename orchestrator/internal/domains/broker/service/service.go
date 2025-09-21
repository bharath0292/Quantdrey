package brokersservice

import (
	"context"

	brokersentity "github.com/bharath0292/quantdrey/internal/domains/broker/entity"
	brokersrepository "github.com/bharath0292/quantdrey/internal/domains/broker/repository"
)

type brokersService struct {
	repo brokersrepository.IBrokersRepository
}

type IBrokersService interface {
	GetBrokersList(ctx context.Context) ([]*brokersentity.Broker, error)
}

func NewUserService(
	repo brokersrepository.IBrokersRepository,
) IBrokersService {
	return &brokersService{repo: repo}
}

func (b *brokersService) GetBrokersList(ctx context.Context) ([]*brokersentity.Broker, error) {
	return b.repo.GetBrokers(ctx)
}
