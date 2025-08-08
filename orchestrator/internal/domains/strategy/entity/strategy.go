package strategyentity

import "time"

type Strategy struct {
	Id                   int          `bson:"id"`
	Name                 string       `bson:"name"`
	UserId               string       `bson:"userId"`
	StartTime            time.Time    `bson:"startTime"`
	EndTime              time.Time    `bson:"endTime"`
	CreatedAt            *time.Time   `bson:"createdAt,omitempty"`
	UpdatedAt            *time.Time   `bson:"updatedAt,omitempty"`
	Rule                 StrategyRule `bson:"rule"`
	MaxTransactionPerDay *int         `bson:"maxTransactionPerDay,omitempty"`
	MaxProfit            *float64     `bson:"maxProfit,omitempty"`
	MaxLoss              *float64     `bson:"maxLoss,omitempty"`
}
