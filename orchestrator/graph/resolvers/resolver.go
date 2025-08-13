package resolver

import (
	graph "github.com/bharath0292/quantdrey/graph/generated"
	strategyservice "github.com/bharath0292/quantdrey/internal/domains/strategy/service"
)

type Resolver struct {
	strategyService strategyservice.IStrategyService
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func NewResolver(strategyService strategyservice.IStrategyService) Resolver {
	return Resolver{
		strategyService: strategyService,
	}
}
