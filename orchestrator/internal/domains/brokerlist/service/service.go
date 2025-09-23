package brokerlistservice

import (
	"context"

	brokerlistentity "github.com/bharath0292/quantdrey/internal/domains/brokerlist/entity"
	brokerlistrepository "github.com/bharath0292/quantdrey/internal/domains/brokerlist/repository"
)

type brokerlistService struct {
	repo brokerlistrepository.IBrokersRepository
}

type IBrokersService interface {
	GetBrokersList(ctx context.Context) ([]*brokerlistentity.Broker, error)
}

func NewUserService(
	repo brokerlistrepository.IBrokersRepository,
) IBrokersService {
	return &brokerlistService{repo: repo}
}

func (b *brokerlistService) GetBrokersList(ctx context.Context) ([]*brokerlistentity.Broker, error) {
	return b.repo.GetBrokers(ctx)
}
