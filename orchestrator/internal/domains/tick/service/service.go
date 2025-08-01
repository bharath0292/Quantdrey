package tickservice

import (
	"encoding/json"
	"time"

	natFactory "github.com/bharath0292/quantdrey/infrastructure/nats"
	tickentity "github.com/bharath0292/quantdrey/internal/domains/tick/entity"
	tickhub "github.com/bharath0292/quantdrey/internal/domains/tick/hub"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

type tickService struct {
	natsClient *natFactory.NatsClient
	tickHub    tickhub.ITickHub
}

type ITickService interface {
	SendRequest(requestMessage *nats.Msg) (bool, error)
	GetTick(time *time.Time, symbol string) *tickentity.Bar
	GetLatestTick(symbol string) *tickentity.Bar
	UpdateTick(tickMessage *nats.Msg)
}

func NewTickService(nc *natFactory.NatsClient, tickHub tickhub.ITickHub) ITickService {
	return &tickService{nc, tickHub}
}

func (th *tickService) SendRequest(requestMessage *nats.Msg) (bool, error) {
	_, err := th.natsClient.SendRequest(requestMessage)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (th *tickService) GetTick(time *time.Time, symbol string) *tickentity.Bar {
	return th.tickHub.FetchLatestTick(symbol)
}

func (th *tickService) GetLatestTick(symbol string) *tickentity.Bar {
	return th.tickHub.FetchLatestTick(symbol)
}

func (th *tickService) UpdateTick(tickMessage *nats.Msg) {
	bar := &tickentity.Bar{}
	if err := json.Unmarshal(tickMessage.Data, bar); err != nil {

		log.Warn().Msgf("Failed to decode bar: %v", err)
		return
	}
	th.tickHub.UpdateTick(tickMessage)
}
