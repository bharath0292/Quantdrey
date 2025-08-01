package indicatorcollection

import (
	"fmt"

	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
)

type Output string

const (
	Value Output = "value"
)

type RsiParams struct {
	Period int
	Output Output
}

type Rsi struct {
	Rsi float64
}

func GetRsi(bar *tickentity.Bar, period int) (*Rsi, bool) {
	key := fmt.Sprintf("rsi_%d", period)

	if data, ok := bar.Indicators[key]; ok {
		return &Rsi{Rsi: data["value"]}, true
	}
	return nil, false
}
