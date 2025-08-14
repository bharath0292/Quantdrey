package strategyentity

import (
	"time"

	"github.com/bharath0292/quantdrey/graph/schema/scalars"
	strategyutils "github.com/bharath0292/quantdrey/internal/domains/strategy/utils"
)

type Strategy struct {
	Id                   scalars.BsonId `bson:"_id,omitempty"`
	Name                 string         `bson:"name"`
	UserId               int            `bson:"userId"`
	StartTime            string         `bson:"startTime"`
	EndTime              string         `bson:"endTime"`
	CreatedAt            *time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt            *time.Time     `bson:"updatedAt,omitempty"`
	Rule                 StrategyRule   `bson:"rule"`
	MaxTransactionPerDay *int           `bson:"maxTransactionPerDay,omitempty"`
	MaxProfit            *float64       `bson:"maxProfit,omitempty"`
	MaxLoss              *float64       `bson:"maxLoss,omitempty"`
}

func (s *Strategy) Validate() error {
	err := strategyutils.ValidateStrategyTime(s.StartTime)
	if err != nil {
		return err
	}
	err = strategyutils.ValidateStrategyTime(s.EndTime)
	if err != nil {
		return err
	}

	return nil
}
