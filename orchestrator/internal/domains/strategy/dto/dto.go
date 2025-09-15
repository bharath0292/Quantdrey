package strategydto

import (
	"fmt"
	"time"

	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateStrategy struct {
	Name                 string
	StartTime            string
	EndTime              string
	Rule                 strategyentity.StrategyRule
	MaxTransactionPerDay *int
	MaxProfit            *float64
	MaxLoss              *float64
}

func (c CreateStrategy) ToEntity() strategyentity.Strategy {
	return strategyentity.Strategy{
		Name:                 c.Name,
		StartTime:            c.StartTime,
		EndTime:              c.EndTime,
		Rule:                 c.Rule,
		MaxTransactionPerDay: c.MaxTransactionPerDay,
		MaxProfit:            c.MaxProfit,
		MaxLoss:              c.MaxLoss,
	}
}

type UpdateStrategy struct {
	Name                 *string
	StartTime            *string
	EndTime              *string
	Rule                 *strategyentity.UpdateStrategyRule
	MaxTransactionPerDay *int
	MaxProfit            *float64
	MaxLoss              *float64
}

func buildExpressionUpdate(prefix string, expr *strategyentity.UpdateExpression, update bson.M) {
	// If Condition is present
	if expr.Condition != nil {
		cond := expr.Condition
		update[prefix+".condition.left"] = cond.Left
		update[prefix+".condition.operator"] = cond.Operator
		update[prefix+".condition.right"] = cond.Right
	}

	// If Logical group is present
	if expr.Logical != nil {
		logical := expr.Logical
		update[prefix+".logical.operator"] = logical.Operator

		for i, subExpr := range logical.Expressions {
			buildExpressionUpdate(fmt.Sprintf("%s.logical.expressions.%d", prefix, i), subExpr, update)
		}
	}
}

func (u UpdateStrategy) ToUpdateBson() bson.M {
	update := bson.M{}

	update["updatedAt"] = time.Now()

	if u.Name != nil {
		update["name"] = *u.Name
	}
	if u.StartTime != nil {
		update["startTime"] = *u.StartTime
	}
	if u.EndTime != nil {
		update["endTime"] = *u.EndTime
	}
	if u.MaxTransactionPerDay != nil {
		update["maxTransactionPerDay"] = *u.MaxTransactionPerDay
	}
	if u.MaxProfit != nil {
		update["maxProfit"] = *u.MaxProfit
	}
	if u.MaxLoss != nil {
		update["maxLoss"] = *u.MaxLoss
	}

	// Handle nested Rule struct
	if u.Rule != nil {

		if u.Rule.Transaction != nil {
			update["rule.transactionType"] = u.Rule.Transaction
		}
		if u.Rule.Profit != nil {
			update["rule.profit"] = u.Rule.Profit
		}
		if u.Rule.Loss != nil {
			update["rule.loss"] = u.Rule.Loss
		}

		// --- SymbolRule ---
		// symbol := u.Rule.Symbols
		// if symbol.LookUpSymbol != nil {
		// 	update["rule.symbol.lookupSymbol"] = symbol.LookUpSymbol
		// }
		// if symbol.Instrument != nil {
		// 	update["rule.symbol.instrumentType"] = symbol.Instrument
		// }
		// if symbol.OrderSymbol != nil {
		// 	update["rule.symbol.orderSymbol"] = symbol.OrderSymbol
		// }
		// if symbol.Expiry != nil {
		// 	update["rule.symbol.expiry"] = symbol.Expiry
		// }
		// if symbol.LotSize != nil {
		// 	update["rule.symbol.lotSize"] = symbol.LotSize
		// }

		// EntryLogic
		if u.Rule.EntryLogic != nil {
			buildExpressionUpdate("rule.entryLogic", u.Rule.EntryLogic, update)
		}

		// ExitLogic
		if u.Rule.ExitLogic != nil {
			buildExpressionUpdate("rule.exitLogic", u.Rule.ExitLogic, update)
		}

	}

	return update
}
