package strategyservice

import (
	"context"
	"errors"
	"log"
	"time"

	strategydto "github.com/bharath0292/quantdrey/internal/domains/strategy/dto"
	strategyentity "github.com/bharath0292/quantdrey/internal/domains/strategy/entity"
	strategyhub "github.com/bharath0292/quantdrey/internal/domains/strategy/hub"
	strategyrepository "github.com/bharath0292/quantdrey/internal/domains/strategy/repository"
	strategyrunner "github.com/bharath0292/quantdrey/internal/domains/strategy/runner"
	tickservice "github.com/bharath0292/quantdrey/internal/domains/tick/service"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type strategyService struct {
	strategyRepository strategyrepository.IStrategyRepository
	tickService        tickservice.ITickService
	hub                strategyhub.IStrategyHub
}

type IStrategyService interface {
	Create(ctx context.Context, userId int, input *strategydto.CreateStrategy) (*strategyentity.Strategy, error)
	Update(ctx context.Context, strategyId bson.ObjectID, input *strategydto.UpdateStrategy) (*strategyentity.Strategy, error)
	Delete(ctx context.Context, strategyId bson.ObjectID) (bool, error)

	Start(ctx context.Context, strategyId bson.ObjectID) (bool, error)
	Stop(ctx context.Context, strategyId bson.ObjectID) error
	Exit(ctx context.Context, strategyId bson.ObjectID) (bool, error)
}

func NewStrategyService(
	strategyRepository strategyrepository.IStrategyRepository,
	tickService tickservice.ITickService,
	hub strategyhub.IStrategyHub,
) IStrategyService {

	return &strategyService{strategyRepository, tickService, hub}
}

func (s *strategyService) Create(ctx context.Context, userId int, input *strategydto.CreateStrategy) (*strategyentity.Strategy, error) {
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

func (s *strategyService) Update(ctx context.Context, strategyId bson.ObjectID, input *strategydto.UpdateStrategy) (*strategyentity.Strategy, error) {
	updateBson := input.ToUpdateBson()

	return s.strategyRepository.UpdateStrategy(ctx, strategyId, updateBson)
}

func (s *strategyService) Delete(ctx context.Context, strategyId bson.ObjectID) (bool, error) {
	return true, nil
}

func (s *strategyService) Start(ctx context.Context, strategyId bson.ObjectID) (bool, error) {
	strategy, err := s.strategyRepository.GetStrategy(ctx, strategyId)
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	runner := strategyrunner.NewStrategyRunner(s.tickService, strategy, cancel)

	s.hub.Register(strategy.Id.Hex(), runner)
	go runner.Start(ctx, strategy)

	return true, nil
}

func (s *strategyService) Stop(ctx context.Context, strategyId bson.ObjectID) error {
	runner := s.hub.Get(strategyId.Hex())
	if runner == nil {
		return errors.New("strategy not found or already stopped")
	}

	runner.Stop()

	s.hub.UnRegister(strategyId.Hex())
	log.Printf("Strategy %s stopped manually", strategyId)
	return nil
}

func (s *strategyService) Exit(ctx context.Context, strategyId bson.ObjectID) (bool, error) {
	return true, nil
}
