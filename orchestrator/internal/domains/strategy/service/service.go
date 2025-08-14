package strategyservice

import (
	"context"
	"time"

	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyrepository "github.com/bharath0292/quantdrey/internal/domains/strategy/repository"
	strategyrunnerservice "github.com/bharath0292/quantdrey/internal/domains/strategyrunner/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type strategyService struct {
	strategyRepository    strategyrepository.IStrategyRepository
	strategyRunnerService strategyrunnerservice.IStrategyRunnerService
}

type IStrategyService interface {
	CreateStrategy(ctx context.Context, userId int, input *strategydto.CreateStrategy) (*strategyentity.Strategy, error)
	UpdateStrategy(ctx context.Context, strategyId bson.ObjectID, input *strategydto.UpdateStrategy) (*strategyentity.Strategy, error)
	RunStrategy(ctx context.Context, strategyId bson.ObjectID) (bool, error)
}

func NewStrategyService(
	strategyRepository strategyrepository.IStrategyRepository,
	strategyRunnerService strategyrunnerservice.IStrategyRunnerService,
) IStrategyService {
	return &strategyService{strategyRepository, strategyRunnerService}
}

func (s *strategyService) CreateStrategy(ctx context.Context, userId int, input *strategydto.CreateStrategy) (*strategyentity.Strategy, error) {
	now := time.Now()
	newStrategy := strategyentity.Strategy{
		Name:                 input.Name,
		StartTime:            input.StartTime,
		EndTime:              input.EndTime,
		Rule:                 input.Rule,
		CreatedAt:            &now,
		MaxTransactionPerDay: input.MaxTransactionPerDay,
		MaxProfit:            input.MaxProfit,
		MaxLoss:              input.MaxLoss,
	}

	err := newStrategy.Validate()
	if err != nil {
		return nil, err
	}

	return s.strategyRepository.CreateStrategy(ctx, &newStrategy)
}

func (s *strategyService) UpdateStrategy(ctx context.Context, strategyId bson.ObjectID, input *strategydto.UpdateStrategy) (*strategyentity.Strategy, error) {
	updateBson := input.ToUpdateBson()

	return s.strategyRepository.UpdateStrategy(ctx, strategyId, updateBson)
}

func (s *strategyService) RunStrategy(ctx context.Context, strategyId bson.ObjectID) (bool, error) {
	strategy, err := s.strategyRepository.GetStrategy(ctx, strategyId)
	if err != nil {
		return false, err
	}
	s.strategyRunnerService.Run(strategy)

	return true, nil
}
