package strategyrunner

import (
	"context"
	"sync"
	"time"

	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyutils "github.com/bharath0292/quantdrey/internal/domains/strategy/utils"
	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
	transactiontypes "github.com/bharath0292/quantdrey/pkg/enums/transactions"

	"github.com/rs/zerolog/log"
)

type Position struct {
	EntryPrice       float64
	ExitPrice        float64
	Quantity         float64
	EntryTransaction transactiontypes.Transaction
	Pnl              float64
}

type strategyRunner struct {
	tickService tickservice.ITickService

	strategy *strategyentity.Strategy

	cancel context.CancelFunc
	mu     sync.Mutex

	dayTxnCount  int
	strategyPnl  float64
	positions    []Position
	openPosition *Position
}

type IStrategyRunner interface {
	Start(ctx context.Context)
	Stop()
}

func NewStrategyRunner(
	tickService tickservice.ITickService,
	strategy *strategyentity.Strategy,
	cancel context.CancelFunc,
) IStrategyRunner {
	return &strategyRunner{
		tickService: tickService,
		strategy:    strategy,
		cancel:      cancel,
		dayTxnCount: 0,
		strategyPnl: 0,
		positions:   []Position{},
	}
}

func (s *strategyRunner) handleTick(bar *tickentity.Bar) {
	endTime := strategyutils.ParseTime(s.strategy.EndTime)

	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	if now.After(*endTime) {
		log.Info().Msgf("End time reached for %s", s.strategy.Id.Hex())
		log.Info().Msg("Place Exit order")
		log.Info().Msg("Update DB")
		log.Info().Msg("Send notification to user")

		if s.openPosition != nil {
			s.positions = append(s.positions, *s.openPosition)
		}

		s.cancel()
		return
	}

	if s.openPosition == nil {
		if s.dayTxnCount > *s.strategy.MaxTransactionPerDay {
			log.Info().Msgf("Daily transaction limit reached for %s", s.strategy.Id.Hex())
			log.Info().Msg("Exit strategy")
			log.Info().Msg("Update DB")

			s.cancel()
			return
		}

		if entryTriggered, err := s.strategy.Rule.EntryLogic.Evaluate(bar); err != nil && entryTriggered {
			log.Info().Msg("Entry Triggered. Placing order")
			log.Info().Msg("Update DB")
			log.Info().Msg("Send notification to user")

			// if placeorder success
			s.dayTxnCount++
		}
	}

	if s.openPosition != nil {
		if exitTriggered, err := s.strategy.Rule.ExitLogic.Evaluate(bar); err != nil && exitTriggered {
			log.Info().Msg("Exit Triggered. exiting order")
		}

		if s.strategy.Rule.Profit != nil && s.openPosition.Pnl >= *s.strategy.MaxProfit {
			log.Info().Msg("Position Max profit reached")
		}

		if s.strategy.Rule.Loss != nil && -s.openPosition.Pnl <= *s.strategy.MaxLoss {
			log.Info().Msg("Position Max profit reached")
		}

		if s.strategy.MaxProfit != nil && s.strategyPnl >= *s.strategy.MaxProfit {
			log.Info().Msg("Strategy Max profit reached")
		}

		if s.strategy.MaxLoss != nil && -s.strategyPnl <= *s.strategy.MaxLoss {
			log.Info().Msg("Strategy Max loss reached")
		}
	}

}

func (s *strategyRunner) Start(ctx context.Context) {

	startTime := strategyutils.ParseTime(s.strategy.StartTime)

	if d := time.Until(*startTime); d > 0 {
		select {
		case <-time.After(d):
		case <-ctx.Done():
			return
		}
	}
	log.Info().Msgf("Strategy %s started", s.strategy.Id.Hex())

	/* Subscribe to symbols */
	// s.tickService.Subscribe(s.strategy.Rule.Symbol.LookUpSymbol, 3, indicators *[]string)

	bar := s.tickService.GetLatestTick(s.strategy.Rule.Symbols[0].LookUpSymbol)
	s.handleTick(bar)
}

func (s *strategyRunner) Stop() {
	s.cancel()
}
