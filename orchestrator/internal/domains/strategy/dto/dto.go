package strategydto

import (
	"time"

	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
)

type NewStrategy struct {
	Name                 string
	UserId               string
	StartTime            time.Time
	EndTime              time.Time
	Rule                 strategyentity.StrategyRule
	MaxTransactionPerDay *int
	MaxProfit            *float64
	MaxLoss              *float64
}

type UpdateStrategy struct {
	StrategyId           string
	StartTime            time.Time
	EndTime              time.Time
	Rule                 strategyentity.StrategyRule
	MaxTransactionPerDay *int
	MaxProfit            *float64
	MaxLoss              *float64
}
