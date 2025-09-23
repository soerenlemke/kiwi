package store

import (
	"errors"
	"log/slog"
	"sync"
)

type InMemoryStore[T any] struct {
	log  slog.Logger
	mu   sync.RWMutex
	data map[string]T
}

func NewInMemoryStore[T any](logger slog.Logger) *InMemoryStore[T] {
	return &InMemoryStore[T]{
		log:  logger,
		data: make(map[string]T),
	}
}

func (s *InMemoryStore[T]) Set(key string, value T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	s.log.Info("Set value in memory store", "key", key)

	return nil
}

func (s *InMemoryStore[T]) Get(key string) (T, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	if !ok {
		var zero T
		return zero, false, nil
	}

	s.log.Info("Retrieved value from memory store", "key", key)
	return value, true, nil
}

func (s *InMemoryStore[T]) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.keyInStore(key) {
		delete(s.data, key)
		s.log.Info("Deleted value from memory store", "key", key)
	}

	return errors.New("key not found")
}

func (s *InMemoryStore[T]) keyInStore(key string) bool {
	if _, ok := s.data[key]; ok {
		return true
	}

	return false
}
