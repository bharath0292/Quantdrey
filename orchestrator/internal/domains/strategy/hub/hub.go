package strategyhub

import (
	"fmt"

	inMemoryFactory "github.com/bharath0292/quantdrey/infrastructure/inMemory"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyutils "github.com/bharath0292/quantdrey/internal/domains/strategy/utils"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"

	"github.com/rs/zerolog/log"
)

type strategyHub struct {
	store       *inMemoryFactory.InMemoryClient[int, []*strategyentity.Strategy]
	tickService tickservice.ITickService
}

type IStrategyHub interface {
	Add(strategy *strategyentity.Strategy) error
}

func NewStrategyHub(
	inMemoryClient *inMemoryFactory.InMemoryClient[int, []*strategyentity.Strategy],
	tickService tickservice.ITickService,
) IStrategyHub {
	return &strategyHub{inMemoryClient, tickService}
}

func (s *strategyHub) Add(strategy *strategyentity.Strategy) error {
	container, ok := s.store.Load(strategy.UserId)
	if !ok {
		container = make([]*strategyentity.Strategy, 10)
	}
	container = append(container, strategy)
	s.store.Store(strategy.UserId, container)

	startTime, err := strategyutils.ParseTime(strategy.StartTime)
	if err != nil {
		return err
	}
	endTime, err := strategyutils.ParseTime(strategy.EndTime)
	if err != nil {
		return err
	}

	strategyutils.SleepUntil(*startTime)

	fmt.Printf("%v", startTime)
	fmt.Printf("%v", endTime)

	bar := s.tickService.GetLatestTick(strategy.Rule.Symbol.LookUpSymbol)
	entryTriggered, err := strategy.Rule.EntryLogic.Evaluate(bar)
	if err != nil {
		log.Error().Err(err).Msg("error evaluating strategy entry")
		return err
	}

	if entryTriggered {
		log.Info().Msg("Entry Triggered. Placing order")
	}

	return nil
}
