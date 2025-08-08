package strategyservice

import (
	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyrepository "github.com/bharath0292/quantdrey/internal/domains/strategy/repository"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
	"github.com/rs/zerolog/log"
)

type strategyService struct {
	tickService        tickservice.ITickService
	strategyRepository strategyrepository.IStrategyRepository
}

type IStrategyService interface {
	CreateStrategy(strat *strategydto.NewStrategy) error
	UpdateStrategy(strat *strategydto.NewStrategy) error
	RunStrategy(strat *strategydto.NewStrategy) error
}

func NewStrategyService(
	tickService tickservice.ITickService,
	strategyRepository strategyrepository.IStrategyRepository,
) IStrategyService {
	return &strategyService{tickService, strategyRepository}
}

func (s *strategyService) CreateStrategy(strat *strategydto.NewStrategy) error {
	return s.strategyRepository.CreateStrategy(strat)
}

func (s *strategyService) UpdateStrategy(strat *strategydto.NewStrategy) error {
	return s.strategyRepository.UpdateStrategy(strat)
}

func (s *strategyService) RunStrategy(strat *strategydto.NewStrategy) error {
	bar := s.tickService.GetLatestTick(strat.Rule.Symbol.LookUpSymbol)
	_, err := strat.Rule.EntryLogic.Evaluate(bar)
	if err != nil {
		log.Error().Err(err).Msg("Error evaluating strategy entry")
		return err
	}
	return nil
}
