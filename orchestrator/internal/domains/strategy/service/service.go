package strategyservice

import (
	"context"
	"time"

	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyrepository "github.com/bharath0292/quantdrey/internal/domains/strategy/repository"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type strategyService struct {
	tickService        tickservice.ITickService
	strategyRepository strategyrepository.IStrategyRepository
}

type IStrategyService interface {
	CreateStrategy(ctx context.Context, userId int, input *strategydto.CreateStrategy) (*strategyentity.Strategy, error)
	UpdateStrategy(ctx context.Context, strategyId bson.ObjectID, input *strategydto.UpdateStrategy) (*strategyentity.Strategy, error)
	RunStrategy(strategyId bson.ObjectID) error
}

func NewStrategyService(
	tickService tickservice.ITickService,
	strategyRepository strategyrepository.IStrategyRepository,
) IStrategyService {
	return &strategyService{tickService, strategyRepository}
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

	return s.strategyRepository.CreateStrategy(ctx, &newStrategy)
}

func (s *strategyService) UpdateStrategy(ctx context.Context, strategyId bson.ObjectID, input *strategydto.UpdateStrategy) (*strategyentity.Strategy, error) {
	updateBson := input.ToUpdateBson()

	return s.strategyRepository.UpdateStrategy(ctx, strategyId, updateBson)
}

func (s *strategyService) RunStrategy(strategyId bson.ObjectID) error {
	// bar := s.tickService.GetLatestTick(strat.Rule.Symbol.LookUpSymbol)
	// _, err := strat.Rule.EntryLogic.Evaluate(bar)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error evaluating strategy entry")
	// 	return err
	// }
	return nil
}
