package resolver

import strategyservice "github.com/bharath0292/quantdrey/internal/domains/strategy/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	strategyService strategyservice.IStrategyService
}

func NewResolver(strategyService strategyservice.IStrategyService) Resolver {
	return Resolver{
		strategyService: strategyService,
	}
}
