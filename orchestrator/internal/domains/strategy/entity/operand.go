package strategyentity

import (
	indicatordto "github.com/bharath0292/quantdrey/internal/domains/indicator/dto"
	"github.com/bharath0292/quantdrey/pkg/enums/indicators"
)

type PriceField string

const (
	PriceFieldClose  PriceField = "close"
	PriceFieldOpen   PriceField = "open"
	PriceFieldHigh   PriceField = "high"
	PriceFieldLow    PriceField = "low"
	PriceFieldVolume PriceField = "volume"
)

type IndicatorField struct {
	Name   indicators.Indicator         `bson:"name"`   // "rsi", "bb", "macd"
	Params indicatordto.IndicatorParams `bson:"params"` // e.g., {"period": 14} or {"period": 20, "stddev": 2}
}

type UpdateIndicatorField struct {
	Name   *indicators.Indicator               `bson:"name"`   // "rsi", "bb", "macd"
	Params *indicatordto.UpdateIndicatorParams `bson:"params"` // e.g., {"period": 14} or {"period": 20, "stddev": 2}
}

type ValueType string

const (
	ValueTypePriceField ValueType = "priceField" // "close", "open", etc.
	ValueTypeIndicator  ValueType = "indicator"  // "rsi", "macd", etc.
	ValueTypeConstant   ValueType = "constant"   // numeric values
)

type Operand struct {
	Type           ValueType       `bson:"type"`
	IndicatorField *IndicatorField `bson:"indicator,omitempty"`  // if Type is indicator
	PriceField     *PriceField     `bson:"priceField,omitempty"` // if Type is price
	ConstantField  *float64        `bson:"constant,omitempty"`   // if Type is constant
	Offset         *int            `bson:"offset,omitempty"`     // e.g., -1 means previous candle
}

type UpdateOperand struct {
	Type           *ValueType            `bson:"type"`
	IndicatorField *UpdateIndicatorField `bson:"indicator,omitempty"`  // if Type is indicator
	PriceField     *PriceField           `bson:"priceField,omitempty"` // if Type is price
	ConstantField  *float64              `bson:"constant,omitempty"`   // if Type is constant
	Offset         *int                  `bson:"offset,omitempty"`     // e.g., -1 means previous candle
}
