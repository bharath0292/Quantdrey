package strategyrepository

import (
	mongoFactory "github.com/bharath0292/quantdrey/infrastructure/mongo"
	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
)

type strategyRepository struct {
	mongo *mongoFactory.MongoClient
}

type IStrategyRepository interface {
	CreateStrategy(strat *strategydto.NewStrategy) error
	UpdateStrategy(strat *strategydto.NewStrategy) error
}

func NewStrategyService(mongo *mongoFactory.MongoClient) IStrategyRepository {
	return &strategyRepository{mongo}
}

func (s *strategyRepository) CreateStrategy(strat *strategydto.NewStrategy) error {
	return s.mongo.CreateDocument()
}

func (s *strategyRepository) UpdateStrategy(strat *strategydto.NewStrategy) error {
	return s.mongo.UpdateDocument()
}
