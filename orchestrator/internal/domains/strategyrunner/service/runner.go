package strategyrunnerservice

import (
	"fmt"

	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyrunnerhub "github.com/bharath0292/quantdrey/internal/domains/strategyrunner/hub"
	strategyrunnerutils "github.com/bharath0292/quantdrey/internal/domains/strategyrunner/utils"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
	"github.com/rs/zerolog/log"
)

type strategyRunnerService struct {
	tickService tickservice.ITickService
	hub         strategyrunnerhub.IStrategyRunnerHub
}

type IStrategyRunnerService interface {
	Run(strategy *strategyentity.Strategy) error
}

func NewStrategyRunnerHub(
	tickService tickservice.ITickService,
	hub strategyrunnerhub.IStrategyRunnerHub,
) IStrategyRunnerService {
	return &strategyRunnerService{tickService, hub}
}

func (s *strategyRunnerService) Run(strategy *strategyentity.Strategy) error {
	startTime := strategyrunnerutils.ParseTime(strategy.StartTime)
	endTime := strategyrunnerutils.ParseTime(strategy.EndTime)

	strategyrunnerutils.SleepUntil(*startTime)

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
