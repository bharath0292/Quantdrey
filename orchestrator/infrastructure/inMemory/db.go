package inMemoryFactory

import (
	"context"

	"github.com/puzpuzpuz/xsync/v4"
)

type InMemoryClient[T any] struct {
	client *xsync.Map[string, T]
}

func NewInMemoryClient[T any](ctx context.Context) (*InMemoryClient[T], error) {
	client := xsync.NewMap[string, T]()

	return &InMemoryClient[T]{client}, nil
}

func (m *InMemoryClient[T]) Load(key string) (value T, ok bool) {
	return m.client.Load(key)
}

func (m *InMemoryClient[T]) Store(symbol string, value T) {
	m.client.Store(symbol, value)
}
