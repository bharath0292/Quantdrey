package strategyservice

import (
	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
	"github.com/rs/zerolog/log"
)

type strategyService struct {
	tickService tickservice.ITickService
}

type IStrategyService interface {
	CreateStrategy(strat *strategydto.NewStrategy) bool
}

func NewStrategyService(tickService tickservice.ITickService) IStrategyService {
	return &strategyService{tickService}
}

func (s *strategyService) CreateStrategy(strat *strategydto.NewStrategy) bool {
	bar := s.tickService.GetLatestTick(strat.Rule.Symbol.LookUpSymbol)
	success, err := strat.Rule.EntryLogic.Evaluate(bar)
	if err != nil {
		log.Error().Err(err).Msg("Error evaluating strategy entry")
		return false
	}
	return success
}
