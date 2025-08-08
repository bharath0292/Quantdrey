package strategyentity

import (
	"fmt"

	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
)

type Expression struct {
	Condition *Condition    `bson:"condition,omitempty" json:"condition,omitempty"`
	Logical   *LogicalGroup `bson:"logical,omitempty" json:"logical,omitempty"`
}

func (e *Expression) Evaluate(bar *tickentity.Bar) (bool, error) {
	if e.Condition != nil {
		return e.Condition.Evaluate(bar)
	} else if e.Logical != nil {
		res, err := e.Logical.Evaluate(bar)
		return res, err
	}
	return false, fmt.Errorf("invalid expression: neither condition nor logical group set")
}
