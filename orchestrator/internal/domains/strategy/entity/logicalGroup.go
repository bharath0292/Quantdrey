package strategyentity

import (
	"fmt"

	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
	"github.com/bharath0292/quantdrey/pkg/enums/operators"
)

type UpdateLogicalGroup struct {
	Operator    *operators.LogicalOperator `bson:"operator"`
	Expressions []*UpdateExpression        `bson:"expressions"`
}

type LogicalGroup struct {
	Operator    operators.LogicalOperator `bson:"operator"`
	Expressions []Expression              `bson:"expressions"`
}

func (e *LogicalGroup) Evaluate(bar *tickentity.Bar) (bool, error) {
	if len(e.Expressions) == 0 {
		return true, nil
	}

	results := make([]bool, len(e.Expressions))
	for i, expr := range e.Expressions {
		result, err := expr.Evaluate(bar)
		if err != nil {
			return false, err
		}
		results[i] = result
	}

	switch e.Operator {
	case operators.LogicalOperatorAnd:
		for _, result := range results {
			if !result {
				return false, nil
			}
		}
		return true, nil
	case operators.LogicalOperatorOr:
		for _, result := range results {
			if result {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, fmt.Errorf("unknown logical operator: %s", e.Operator)
	}
}
