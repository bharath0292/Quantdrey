package natsfactory

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	natsPool "github.com/octu0/nats-pool"
	"github.com/rs/zerolog/log"
)

type NatsConfig struct {
	URL      string
	PoolSize int
	Options  []nats.Option
}

type NatsClient struct {
	client         *natsPool.ConnPool
	subscriptions  []*nats.Subscription
	subscriptionsM sync.Mutex
}

type MessageHandler func(msg *nats.Msg)

func NewNatsClient(config NatsConfig) (*NatsClient, error) {
	// Apply default options if none are provided
	opts := config.Options
	if opts == nil {
		opts = []nats.Option{nats.NoEcho(), nats.Name("default-client/1.0")}
	}

	// Create pool
	connPool := natsPool.New(config.PoolSize, config.URL, opts...)

	// Validate connectivity
	nc, err := connPool.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get NATS connection from pool: %w", err)
	}
	defer connPool.Put(nc)

	// Perform a test ping
	if err := nc.Publish("PING", []byte("TEST PING FROM CLIENT")); err != nil {
		return nil, fmt.Errorf("failed to publish test ping: %w", err)
	}

	return &NatsClient{
		client: connPool,
	}, nil
}

func (nc *NatsClient) Close() {
	nc.client.DisconnectAll()
}

func (nc *NatsClient) GetClient() (*nats.Conn, error) {
	return nc.client.Get()
}

func (nc *NatsClient) Subscribe(subjectHandlers map[string]MessageHandler) {
	conn, err := nc.client.Get()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get NATS connection from pool")
		return
	}
	defer nc.client.Put(conn)

	for subject, handler := range subjectHandlers {
		subjectCopy := subject
		handlerCopy := handler

		sub, err := conn.Subscribe(subjectCopy, func(msg *nats.Msg) {
			if handlerCopy != nil {
				handlerCopy(msg)
			} else {
				log.Warn().Msgf("No handler defined for subject: %s", msg.Subject)
			}
		})
		if err != nil {
			log.Error().Err(err).Msgf("Failed to subscribe to subject: %s", subjectCopy)
			continue
		}

		nc.subscriptionsM.Lock()
		nc.subscriptions = append(nc.subscriptions, sub)
		nc.subscriptionsM.Unlock()

		log.Info().Msgf("Subscribed to subject: %s", subjectCopy)
	}
}

func (nc *NatsClient) UnsubscribeAll() {
	nc.subscriptionsM.Lock()
	defer nc.subscriptionsM.Unlock()

	for _, sub := range nc.subscriptions {
		if err := sub.Unsubscribe(); err != nil {
			log.Error().Err(err).Msgf("Failed to unsubscribe from subject: %s", sub.Subject)
		} else {
			log.Info().Msgf("Unsubscribed from subject: %s", sub.Subject)
		}
	}
	nc.subscriptions = nil
}

func (nc *NatsClient) SendRequest(requestMessage *nats.Msg) ([]byte, error) {
	// Validate required fields
	if requestMessage.Subject == "" {
		return []byte{}, fmt.Errorf("message subject is required")
	}

	conn, err := nc.client.Get()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get NATS connection from pool: %w", err)
	}
	defer nc.client.Put(conn)

	// Send a request and wait up to 10 seconds for a response
	response, err := conn.RequestMsg(requestMessage, 10*time.Second)
	if err != nil {
		log.Error().Err(err).Msgf("Error sending request to subject: %s", requestMessage.Subject)
		return []byte{}, err
	}

	log.Debug().Msgf("Received response from subject %s: %s", requestMessage.Subject, string(response.Data))
	return response.Data, nil
}

func (nc *NatsClient) SendMessage(message nats.Msg) error {
	// Validate required fields
	if message.Subject == "" {
		return fmt.Errorf("message subject is required")
	}

	conn, err := nc.client.Get()
	if err != nil {
		return fmt.Errorf("failed to get NATS connection from pool: %w", err)
	}
	defer nc.client.Put(conn)

	if err := conn.PublishMsg(&message); err != nil {
		log.Error().Err(err).Msgf("Failed to publish message to subject: %s", message.Subject)
		return err
	}

	log.Debug().Msgf("Message published to subject: %s", message.Subject)
	return nil
}
