package strategyrepository

import (
	"context"

	mongoFactory "github.com/bharath0292/quantdrey/infrastructure/mongo"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const COLLECTION_NAME = "strategies"

type strategyRepository struct {
	mongo *mongoFactory.MongoClient
}

type IStrategyRepository interface {
	GetStrategy(ctx context.Context, strategyID bson.ObjectID) (*strategyentity.Strategy, error)
	CreateStrategy(ctx context.Context, newStrategy *strategyentity.Strategy) (*strategyentity.Strategy, error)
	UpdateStrategy(ctx context.Context, strategyID bson.ObjectID, updateMap bson.M) (*strategyentity.Strategy, error)
}

func NewStrategyService(mongo *mongoFactory.MongoClient) IStrategyRepository {
	return &strategyRepository{mongo}
}

func (s *strategyRepository) GetStrategy(ctx context.Context, strategyID bson.ObjectID) (*strategyentity.Strategy, error) {
	var strategy strategyentity.Strategy
	err := s.mongo.ReadDocument(ctx, COLLECTION_NAME, strategyID, &strategy)
	if err != nil {
		return nil, err
	}

	return &strategy, nil
}

func (s *strategyRepository) CreateStrategy(ctx context.Context, newStrategy *strategyentity.Strategy) (*strategyentity.Strategy, error) {
	var createdStrategy strategyentity.Strategy
	_, err := s.mongo.CreateDocument(ctx, COLLECTION_NAME, newStrategy, &createdStrategy)
	if err != nil {
		return nil, err
	}

	return &createdStrategy, nil
}

func (s *strategyRepository) UpdateStrategy(ctx context.Context, strategyID bson.ObjectID, updateMap bson.M) (*strategyentity.Strategy, error) {
	var updatedStrategy strategyentity.Strategy

	_, err := s.mongo.UpdateDocument(
		ctx,
		COLLECTION_NAME,
		bson.M{"_id": strategyID},
		bson.M{"$set": updateMap},
		&updatedStrategy,
	)
	if err != nil {
		return nil, err
	}

	return &updatedStrategy, nil
}
