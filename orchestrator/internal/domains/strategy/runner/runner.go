package strategyrunner

import (
	"context"
	"time"

	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyutils "github.com/bharath0292/quantdrey/internal/domains/strategy/utils"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"

	"github.com/rs/zerolog/log"
)

type strategyRunner struct {
	tickService tickservice.ITickService

	strategy *strategyentity.Strategy
	cancel   context.CancelFunc
	// mu       sync.Mutex

	// currentProfit float64
	// currentLoss   float64
}

type IStrategyRunner interface {
	Start(ctx context.Context, strategy *strategyentity.Strategy) error
	Stop()
}

func NewStrategyRunner(tickService tickservice.ITickService, strategy *strategyentity.Strategy, cancel context.CancelFunc) IStrategyRunner {
	return &strategyRunner{tickService, strategy, cancel}
}

func (s *strategyRunner) Stop() {
	s.cancel()
}

func (s *strategyRunner) Start(ctx context.Context, strategy *strategyentity.Strategy) error {
	startTime := strategyutils.ParseTime(strategy.StartTime)
	endTime := strategyutils.ParseTime(strategy.EndTime)

	// To subscribe to symbols before 10 seconds of strategy starts
	symbolSubscribeTime := startTime.Add(-10 * time.Second)
	strategyutils.SleepUntil(symbolSubscribeTime)
	/* Subscribe to symbols */
	// s.tickService.Subscribe(strategy.Rule.Symbol.LookUpSymbol, timeframe, indicators *[]string)

	strategyutils.SleepUntil(*startTime)

	bar := s.tickService.GetLatestTick(strategy.Rule.Symbol.LookUpSymbol)
	entryTriggered, err := strategy.Rule.EntryLogic.Evaluate(bar)
	if err != nil {
		log.Error().Err(err).Msg("error evaluating strategy entry")
		return err
	}

	if entryTriggered {
		log.Info().Msg("Entry Triggered. Placing order")
	}

	if time.Now().After(*endTime) {
		log.Info().Msg("Time reached closing order")
	}

	return nil
}
