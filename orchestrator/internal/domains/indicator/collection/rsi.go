package indicatorcollection

import (
	"fmt"

	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
)

type RsiOutput string

const (
	RsiOutputValue RsiOutput = "rsi"
)

type RsiParams struct {
	Period int       `bson:"period"`
	Output RsiOutput `bson:"output"`
}

type UpdateRsiParams struct {
	Period *int       `bson:"period"`
	Output *RsiOutput `bson:"output"`
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
