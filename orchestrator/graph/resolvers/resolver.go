package resolver

import (
	graph "github.com/bharath0292/quantdrey/graph/generated"
	brokersservice "github.com/bharath0292/quantdrey/internal/domains/broker/service"
	strategyservice "github.com/bharath0292/quantdrey/internal/domains/strategy/service"
	userservice "github.com/bharath0292/quantdrey/internal/domains/user/service"
)

type Resolver struct {
	userService     userservice.IUserService
	brokersService  brokersservice.IBrokersService
	strategyService strategyservice.IStrategyService
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func NewResolver(userService userservice.IUserService, brokersService brokersservice.IBrokersService, strategyService strategyservice.IStrategyService) Resolver {
	return Resolver{
		userService:     userService,
		brokersService:  brokersService,
		strategyService: strategyService,
	}
}
