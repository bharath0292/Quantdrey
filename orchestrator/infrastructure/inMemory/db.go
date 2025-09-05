package inMemoryFactory

import (
	"github.com/puzpuzpuz/xsync/v4"
)

type InMemoryClient[K comparable, V any] struct {
	client *xsync.Map[K, V]
}

func NewInMemoryClient[K comparable, V any]() (*InMemoryClient[K, V], error) {
	client := xsync.NewMap[K, V]()

	return &InMemoryClient[K, V]{client}, nil
}

func (m *InMemoryClient[K, V]) Load(key K) (value V, ok bool) {
	return m.client.Load(key)
}

func (m *InMemoryClient[K, V]) Store(key K, value V) {
	m.client.Store(key, value)
}

func (m *InMemoryClient[K, V]) Delete(key K) {
	m.client.Delete(key)
}
