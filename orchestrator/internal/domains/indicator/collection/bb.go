package indicatorcollection

import (
	"fmt"

	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
)

type BbOutput string

const (
	Upper  BbOutput = "upper"
	Middle BbOutput = "middle"
	Lower  BbOutput = "lower"
)

type BbParams struct {
	Period int
	StdDev int
	Output BbOutput
}

type Bb struct {
	Upper  float64 `json:"upper"`
	Middle float64 `json:"middle"`
	Lower  float64 `json:"lower"`
}

func GetBb(bar *tickentity.Bar, period, stddev int) (*Bb, bool) {
	key := fmt.Sprintf("bb_%d_%d", period, stddev)
	if data, ok := bar.Indicators[key]; ok {

		return &Bb{
			Upper:  data["upper"],
			Middle: data["middle"],
			Lower:  data["lower"],
		}, true
	}
	return nil, false
}
