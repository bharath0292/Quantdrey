package strategyentity

import (
	instrumenttype "github.com/bharath0292/quantdrey/pkg/enums/instruments"
	transactiontype "github.com/bharath0292/quantdrey/pkg/enums/transactions"
)

type UpdateSymbolRule struct {
	LookUpSymbol *string                    `bson:"lookupSymbol"`
	Instrument   *instrumenttype.Instrument `bson:"instrumentType"`
	OrderSymbol  *string                    `bson:"orderSymbol"`
	Expiry       *string                    `bson:"expiry,omitempty"`
	LotSize      *int                       `bson:"lotSize"`
}

type SymbolRule struct {
	LookUpSymbol string                    `bson:"lookupSymbol"`
	Instrument   instrumenttype.Instrument `bson:"instrumentType"`
	OrderSymbol  string                    `bson:"orderSymbol"`
	Expiry       string                    `bson:"expiry,omitempty"`
	LotSize      int                       `bson:"lotSize"`
}

type UpdateStrategyRule struct {
	Symbol      *UpdateSymbolRule            `bson:"symbol"`
	Transaction *transactiontype.Transaction `bson:"transactionType"`
	Profit      *float64                     `bson:"profit,omitempty"`
	Loss        *float64                     `bson:"loss,omitempty"`
	EntryLogic  *UpdateExpression            `bson:"entryLogic"`
	ExitLogic   *UpdateExpression            `bson:"exitLogic"`
}

type StrategyRule struct {
	Symbol      SymbolRule                  `bson:"symbol"`
	Transaction transactiontype.Transaction `bson:"transactionType"`
	Profit      *float64                    `bson:"profit,omitempty"`
	Loss        *float64                    `bson:"loss,omitempty"`
	EntryLogic  Expression                  `bson:"entryLogic"`
	ExitLogic   Expression                  `bson:"exitLogic"`
}
