package strategyrepository

import (
	"context"
	"fmt"

	mongoFactory "github.com/bharath0292/quantdrey/infrastructure/mongo"
	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
)

const COLLECTION_NAME = "strategies"

type strategyRepository struct {
	mongo *mongoFactory.MongoClient
}

type IStrategyRepository interface {
	CreateStrategy(ctx context.Context, newStrategy *strategyentity.Strategy) (*strategyentity.Strategy, error)
	UpdateStrategy(ctx context.Context, strat *strategydto.NewStrategy) error
}

func NewStrategyService(mongo *mongoFactory.MongoClient) IStrategyRepository {
	return &strategyRepository{mongo}
}

func (s *strategyRepository) CreateStrategy(ctx context.Context, newStrategy *strategyentity.Strategy) (*strategyentity.Strategy, error) {
	var createdStrategy strategyentity.Strategy
	_, err := s.mongo.CreateDocument(ctx, COLLECTION_NAME, newStrategy, &createdStrategy)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", createdStrategy)

	return &createdStrategy, nil
}

func (s *strategyRepository) UpdateStrategy(ctx context.Context, strat *strategydto.NewStrategy) error {
	return s.mongo.UpdateDocument()
}
