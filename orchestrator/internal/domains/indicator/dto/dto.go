package indicatordto

import (
	indicatorcollection "github.com/bharath0292/quantdrey/internal/domains/indicator/collection"
)

type IndicatorParams struct {
	Rsi *indicatorcollection.RsiParams
	Bb  *indicatorcollection.BbParams
}
