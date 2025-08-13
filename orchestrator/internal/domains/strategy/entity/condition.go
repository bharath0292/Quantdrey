package strategyentity

import (
	"fmt"

	indicatorcollection "github.com/bharath0292/quantdrey/internal/domains/indicator/collection"
	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
	"github.com/bharath0292/quantdrey/pkg/enums/indicators"
	"github.com/bharath0292/quantdrey/pkg/enums/operators"
)

type UpdateCondition struct {
	Left     *UpdateOperand                `bson:"left"`
	Operator *operators.ComparisonOperator `bson:"operator"`
	Right    *UpdateOperand                `bson:"right"`
}

type Condition struct {
	Left     Operand                      `bson:"left"`
	Operator operators.ComparisonOperator `bson:"operator"`
	Right    Operand                      `bson:"right"`
}

func (e *Condition) compareValues(left *float64, op operators.ComparisonOperator, right *float64) bool {
	switch op {
	case operators.ComparisonOperatorEqual:
		return left == right
	case operators.ComparisonOperatorNotEqual:
		return left != right
	case operators.ComparisonOperatorGreater:
		return *left > *right
	case operators.ComparisonOperatorGreaterOrEqual:
		return *left >= *right
	case operators.ComparisonOperatorLess:
		return *left < *right
	case operators.ComparisonOperatorLessOrEqual:
		return *left <= *right
	default:
		return false
	}
}

func (e *Condition) getOperandValue(operand Operand, bar *tickentity.Bar) (*float64, error) {
	switch operand.Type {
	case ValueTypeConstant:
		if operand.ConstantField == nil {
			return nil, fmt.Errorf("empty constant field")
		}
		return operand.ConstantField, nil

	case ValueTypePriceField:
		if operand.PriceField == nil {
			return nil, fmt.Errorf("empty price field")
		}
		switch *operand.PriceField {
		case PriceFieldOpen:
			return &bar.Open, nil
		case PriceFieldHigh:
			return &bar.High, nil
		case PriceFieldLow:
			return &bar.Low, nil
		case PriceFieldClose:
			return &bar.Close, nil
		case PriceFieldVolume:
			return &bar.Volume, nil
		default:
			return nil, fmt.Errorf("unknown price field: %s", *operand.PriceField)
		}

	case ValueTypeIndicator:
		if operand.IndicatorField == nil {
			return nil, fmt.Errorf("empty indicator field")
		}

		switch operand.IndicatorField.Name {
		case indicators.IndicatorRsi:
			indicator, isAvailable := indicatorcollection.GetRsi(bar, operand.IndicatorField.Params.Rsi.Period)
			if !isAvailable {
				return nil, fmt.Errorf("indicator not available")
			}
			return &indicator.Rsi, nil
		}

		return nil, fmt.Errorf("unknown indicator field: %s", operand.IndicatorField.Name)

	default:
		return nil, fmt.Errorf("unknown operand type: %s", operand.Type)
	}
}

func (e *Condition) Evaluate(bar *tickentity.Bar) (bool, error) {
	leftVal, err := e.getOperandValue(e.Left, bar)
	if err != nil {
		return false, err
	}

	rightVal, err := e.getOperandValue(e.Right, bar)
	if err != nil {
		return false, err
	}

	return e.compareValues(leftVal, e.Operator, rightVal), nil
}
