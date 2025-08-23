package strategyhub

import (
	inMemoryFactory "github.com/bharath0292/quantdrey/infrastructure/inMemory"
	strategyrunner "github.com/bharath0292/quantdrey/internal/domains/strategy/runner"
)

type strategyHub struct {
	store *inMemoryFactory.InMemoryClient[string, strategyrunner.IStrategyRunner]
}

type IStrategyHub interface {
	Register(id string, strategy strategyrunner.IStrategyRunner)
	UnRegister(id string)
	Get(id string) strategyrunner.IStrategyRunner
}

func NewStrategyHub(
	inMemoryClient *inMemoryFactory.InMemoryClient[string, strategyrunner.IStrategyRunner],
) IStrategyHub {
	return &strategyHub{inMemoryClient}
}

func (s *strategyHub) Register(id string, strategy strategyrunner.IStrategyRunner) {
	s.store.Store(id, strategy)
}

func (s *strategyHub) UnRegister(id string) {
	s.store.Delete(id)
}

func (s *strategyHub) Get(id string) strategyrunner.IStrategyRunner {
	strategy, found := s.store.Load(id)
	if !found {
		return nil
	}
	return strategy
}
