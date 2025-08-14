package resolver

import (
	"context"
	"fmt"

	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *mutationResolver) CreateStrategy(ctx context.Context, userID int, input strategydto.CreateStrategy) (*strategyentity.Strategy, error) {

	createdStrategy, err := r.strategyService.CreateStrategy(ctx, userID, &input)
	if err != nil {
		return nil, err
	}

	return createdStrategy, nil
}

func (r *mutationResolver) UpdateStrategy(ctx context.Context, strategyID bson.ObjectID, input strategydto.UpdateStrategy) (*strategyentity.Strategy, error) {

	updatedStrategy, err := r.strategyService.UpdateStrategy(ctx, strategyID, &input)
	if err != nil {
		return nil, err
	}

	return updatedStrategy, nil
}

func (r *mutationResolver) RunStrategy(ctx context.Context, strategyID bson.ObjectID) (bool, error) {
	return r.strategyService.RunStrategy(ctx, strategyID)
}

func (r *queryResolver) ListAllStrategy(ctx context.Context, userID int) (*strategyentity.Strategy, error) {
	panic(fmt.Errorf("not implemented: ListAllStrategy - listAllStrategy"))
}
