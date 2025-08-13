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
	Period int      `bson:"period"`
	StdDev int      `bson:"stddev"`
	Output BbOutput `bson:"output"`
}

type UpdateBbParams struct {
	Period *int      `bson:"period"`
	StdDev *int      `bson:"stddev"`
	Output *BbOutput `bson:"output"`
}

type Bb struct {
	Upper  float64
	Middle float64
	Lower  float64
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
