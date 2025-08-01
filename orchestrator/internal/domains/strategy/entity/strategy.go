package strategyentity

import "time"

type Strategy struct {
	Id                   int          `bson:"id"`
	Name                 string       `bson:"name"`
	UserId               string       `bson:"user_id"`
	StartTime            time.Time    `bson:"start_time"`
	EndTime              time.Time    `bson:"end_time"`
	CreatedAt            *time.Time   `bson:"created_at,omitempty"`
	UpdatedAt            *time.Time   `bson:"updated_at,omitempty"`
	Rule                 StrategyRule `bson:"rule"`
	MaxTransactionPerDay *int         `bson:"max_transaction_per_day,omitempty"`
	MaxProfit            *float64     `bson:"max_profit,omitempty"`
	MaxLoss              *float64     `bson:"max_loss,omitempty"`
}
