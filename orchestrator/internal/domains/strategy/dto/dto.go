package strategydto

import (
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
)

type NewStrategy struct {
	Name                 string
	StartTime            string
	EndTime              string
	Rule                 strategyentity.StrategyRule
	MaxTransactionPerDay *int
	MaxProfit            *float64
	MaxLoss              *float64
}

type UpdateStrategy struct {
	StrategyId           *string
	StartTime            *string
	EndTime              *string
	Rule                 *strategyentity.StrategyRule
	MaxTransactionPerDay *int
	MaxProfit            *float64
	MaxLoss              *float64
}
