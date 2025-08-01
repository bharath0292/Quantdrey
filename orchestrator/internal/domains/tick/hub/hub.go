package tickhub

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"

	natsfactory "github.com/bharath0292/quantdrey/infrastructure/nats"
	redisFactory "github.com/bharath0292/quantdrey/infrastructure/redis"
	tickdto "github.com/bharath0292/quantdrey/internal/domains/tick/dto"
	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"

	"github.com/nats-io/nats.go"
	"github.com/puzpuzpuz/xsync/v4"
	"github.com/rs/zerolog/log"
)

type tickHub struct {
	inMemoryClient struct {
		client *xsync.Map[string, tickentity.Bar]
	}
	redisClient *redisFactory.RedisClient
	natsClient  *natsfactory.NatsClient
}

type ITickHub interface {
	SubscribeTicks()
	Register(symbol string, timeframe uint8, indicators *[]string) error
	UnRegister(symbol string, timeframe uint8) error
	FetchLatestTick(symbol string) *tickentity.Bar
	UpdateTick(tickMessage *nats.Msg)
}

func NewTickHub(redisClient *redisFactory.RedisClient, natsClient *natsfactory.NatsClient) ITickHub {
	store := xsync.NewMap[string, tickentity.Bar]()

	return &tickHub{
		struct {
			client *xsync.Map[string, tickentity.Bar]
		}{store},
		redisClient,
		natsClient,
	}
}

func (th *tickHub) SubscribeTicks() {
	go func() {
		th.natsClient.Subscribe(map[string]natsfactory.MessageHandler{
			"tick": th.UpdateTick,
		})
	}()
}

func (th *tickHub) Register(symbol string, timeframe uint8, indicators *[]string) error {
	key := fmt.Sprintf("orchestrator:subscribed:%s:%d", symbol, timeframe)
	var newIndicatorsMap map[string]struct{}
	var newIndicatorsList []string

	if indicators != nil && len(*indicators) > 0 {
		newIndicatorsMap = make(map[string]struct{})
		for _, i := range *indicators {
			newIndicatorsMap[i] = struct{}{}
		}
		newIndicatorsList = *indicators
	}

	keyExists, err := th.redisClient.KeyExists(key)
	if err != nil {
		return err
	}

	missingIndicators := []string{}
	if keyExists {
		current, err := th.redisClient.GetSetMembers(key)
		if err != nil {
			return err
		}

		currentMap := make(map[string]struct{})
		for _, c := range current {
			currentMap[c] = struct{}{}
		}

		for ind := range newIndicatorsMap {
			if _, ok := currentMap[ind]; !ok {
				missingIndicators = append(missingIndicators, ind)
			}
		}
	} else {
		// First time requesting this symbol+timeframe → all indicators are "missing"
		missingIndicators = newIndicatorsList
	}

	// If it's a new symbol+timeframe and there are no indicators, still send a request for bar-only
	if !keyExists || len(missingIndicators) > 0 {
		// Only send missing indicators (or nil if none)
		req, err := tickdto.NewBarRequest(tickdto.Subscribe, symbol, timeframe, missingIndicators)
		if err != nil {
			return fmt.Errorf("failed to build request: %w", err)
		}

		_, err = th.natsClient.SendRequest(&nats.Msg{
			Subject: "orchestrator.subscribe",
			Data:    req,
		})
		if err != nil {
			return fmt.Errorf("failed to send subscription request: %w", err)
		}
	}

	if len(missingIndicators) > 0 {
		if err := th.redisClient.AddToSet(key, missingIndicators...); err != nil {
			return fmt.Errorf("failed to update redis set: %w", err)
		}
	} else if !keyExists {
		// First-time request with no indicators: create empty set to track bar-only request
		if err := th.redisClient.AddToSet(key); err != nil {
			return fmt.Errorf("failed to initialize redis set: %w", err)
		}
	}

	return nil
}

func (th *tickHub) UnRegister(symbol string, timeframe uint8) error {
	key := fmt.Sprintf("orchestrator:subscribed:%s:%d", symbol, timeframe)
	countKey := key + ":count"

	count, err := th.redisClient.DecrementCount(countKey)
	if err != nil {
		return fmt.Errorf("failed to decrement subscription count: %w", err)
	}

	if count > 0 {
		// Still in use by other users
		return nil
	}

	// Fully unsubscribe
	req, err := tickdto.NewBarRequest(tickdto.Unsubscribe, symbol, timeframe, nil)
	if err != nil {
		return fmt.Errorf("failed to build unsubscribe request: %w", err)
	}

	_, err = th.natsClient.SendRequest(&nats.Msg{
		Subject: "orchestrator.unsubscribe",
		Data:    req,
	})
	if err != nil {
		return fmt.Errorf("failed to send unsubscribe request: %w", err)
	}

	// Cleanup Redis keys
	if err := th.redisClient.DeleteKeys(key, countKey); err != nil {
		return fmt.Errorf("failed to delete redis keys: %w", err)
	}

	return nil
}

func (th *tickHub) FetchTick(time *time.Time, symbol string) *tickentity.Bar {
	return nil
}

func (th *tickHub) FetchLatestTick(symbol string) *tickentity.Bar {
	indicators := make(map[string]map[string]float64)

	// Initialize the inner map before assignment
	indicators["rsi_14"] = make(map[string]float64)
	indicators["rsi_14"]["value"] = 75

	bar := tickentity.Bar{
		Symbol:     "SBIN",
		Timestamp:  123466383,
		Open:       800,
		High:       810,
		Low:        790,
		Close:      805,
		Volume:     123666,
		Indicators: indicators,
	}
	// bar, exists := th.inMemoryClient.client.Load(symbol)
	// if !exists {
	// 	return nil
	// }
	return &bar
}

func (th *tickHub) UpdateTick(tickMessage *nats.Msg) {
	bar := &tickentity.Bar{}
	if err := json.Unmarshal(tickMessage.Data, bar); err != nil {
		log.Warn().Msgf("Failed to decode bar: %v", err)
		return
	}

	th.inMemoryClient.client.Store(bar.Symbol, *bar)
}
