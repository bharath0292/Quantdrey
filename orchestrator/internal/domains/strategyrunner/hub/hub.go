package strategyrunnerhub

import (
	inMemoryFactory "github.com/bharath0292/quantdrey/infrastructure/inMemory"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
)

type strategyRunnerHub struct {
	store       *inMemoryFactory.InMemoryClient[int, []*strategyentity.Strategy]
	tickService tickservice.ITickService
}

type IStrategyRunnerHub interface {
	Add(strategy *strategyentity.Strategy) error
}

func NewStrategyRunnerHub(
	inMemoryClient *inMemoryFactory.InMemoryClient[int, []*strategyentity.Strategy],
	tickService tickservice.ITickService,
) IStrategyRunnerHub {
	return &strategyRunnerHub{inMemoryClient, tickService}
}

func (s *strategyRunnerHub) Add(strategy *strategyentity.Strategy) error {
	container, ok := s.store.Load(strategy.UserId)
	if !ok {
		container = make([]*strategyentity.Strategy, 10)
	}
	container = append(container, strategy)
	s.store.Store(strategy.UserId, container)

	return nil
}
