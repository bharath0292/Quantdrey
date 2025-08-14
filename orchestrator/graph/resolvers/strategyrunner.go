package resolver

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *queryResolver) RunStrategy(ctx context.Context, strategyID bson.ObjectID) (bool, error) {
	return r.strategyService.RunStrategy(ctx, strategyID)
}
