package strategydto

// func toEntitySymbolRule(input SymbolRule) SymbolRule {
// 	return SymbolRule{
// 		LookUpSymbol: input.LookUpSymbol,
// 		Instrument:   input.Instrument,
// 		OrderSymbol:  input.OrderSymbol,
// 		Expiry:       input.Expiry,
// 		LotSize:      input.LotSize,
// 	}
// }

// func toEntityStrategyRule(input StrategyRule) StrategyRule {
// 	return StrategyRule{
// 		Symbol:      toEntitySymbolRule(input.Symbol),
// 		Transaction: input.Transaction,
// 		Profit:      input.Profit,
// 		Loss:        input.Loss,
// 	}
// }

// func NewStrategyToEntity(input NewStrategy) *Strategy {
// 	rules := make([]StrategyRule, len(input.Rule))
// 	for i, rule := range input.Rule {
// 		rules[i] = toEntityStrategyRule(rule)
// 	}

// 	return &Strategy{
// 		Name:                 input.Name,
// 		UserId:               input.UserID,
// 		StartTime:            input.StartTime,
// 		EndTime:              input.EndTime,
// 		Rule:                 rules,
// 		MaxTransactionPerDay: input.MaxTransactionPerDay,
// 		MaxProfit:            input.MaxProfit,
// 		MaxLoss:              input.MaxLoss,
// 	}
// }
