package strategyentity

import (
	instrumenttype "github.com/bharath0292/quantdrey/pkg/enums/instruments"
	transactiontype "github.com/bharath0292/quantdrey/pkg/enums/transactions"
)

type SymbolRule struct {
	LookUpSymbol string                    `bson:"lookup_symbol"`
	Instrument   instrumenttype.Instrument `bson:"instrument_type"`
	OrderSymbol  string                    `bson:"order_symbol"`
	Expiry       string                    `bson:"expiry,omitempty"`
	LotSize      int                       `bson:"lot_size"`
}

type StrategyRule struct {
	Symbol      SymbolRule                  `bson:"symbol"`
	Transaction transactiontype.Transaction `bson:"transaction_type"`
	Profit      float64                     `bson:"max_profit"`
	Loss        float64                     `bson:"max_loss"`
	EntryLogic  *Expression                 `bson:"entry_logic,omitempty"`
	ExitLogic   *Expression                 `bson:"exit_logic,omitempty"`
}
