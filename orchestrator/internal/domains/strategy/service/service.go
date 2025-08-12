package strategyservice

import (
	"context"
	"time"

	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyrepository "github.com/bharath0292/quantdrey/internal/domains/strategy/repository"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
)

type strategyService struct {
	tickService        tickservice.ITickService
	strategyRepository strategyrepository.IStrategyRepository
}

type IStrategyService interface {
	CreateStrategy(ctx context.Context, userId int, input *strategydto.NewStrategy) (*strategyentity.Strategy, error)
	UpdateStrategy(ctx context.Context, input *strategydto.NewStrategy) error
	RunStrategy(strategyId int) error
}

func NewStrategyService(
	tickService tickservice.ITickService,
	strategyRepository strategyrepository.IStrategyRepository,
) IStrategyService {
	return &strategyService{tickService, strategyRepository}
}

func (s *strategyService) CreateStrategy(ctx context.Context, userId int, input *strategydto.NewStrategy) (*strategyentity.Strategy, error) {
	now := time.Now()
	newStrategy := strategyentity.Strategy{
		Name:      input.Name,
		UserId:    userId,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Rule:      input.Rule,
		CreatedAt: &now,
	}
	return s.strategyRepository.CreateStrategy(ctx, &newStrategy)
}

func (s *strategyService) UpdateStrategy(ctx context.Context, strat *strategydto.NewStrategy) error {
	return s.strategyRepository.UpdateStrategy(ctx, strat)
}

func (s *strategyService) RunStrategy(strategyId int) error {
	// bar := s.tickService.GetLatestTick(strat.Rule.Symbol.LookUpSymbol)
	// _, err := strat.Rule.EntryLogic.Evaluate(bar)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error evaluating strategy entry")
	// 	return err
	// }
	return nil
}
