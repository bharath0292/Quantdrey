package strategyhub

import strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"

type strategyRunner struct {
	strategy *strategyentity.Strategy
}

func (s *strategyHub) Run() error {
	return nil
}
