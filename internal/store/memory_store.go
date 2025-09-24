package store

import (
	"errors"
	"kiwi/internal/domain"
	"kiwi/internal/recovery"
	"log/slog"
	"sync"
)

type InMemoryStore[T domain.AllowedTypes] struct {
	log  slog.Logger
	rlw  recovery.RecoverLogWriter[T]
	mu   sync.RWMutex
	data map[string]T
}

func NewInMemoryStore[T domain.AllowedTypes](logger slog.Logger, rlw recovery.RecoverLogWriter[T]) *InMemoryStore[T] {
	return &InMemoryStore[T]{
		log:  logger,
		rlw:  rlw,
		data: make(map[string]T),
	}
}

func (s *InMemoryStore[T]) Set(key string, value T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	s.logSetAction("set", key, value)
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
		s.rlw.LogDeleteAction("delete", key)
		s.log.Info("Deleted value from memory store", "key", key)

		return nil
	} else {
		return errors.New("key not found")
	}
}

func (s *InMemoryStore[T]) keyInStore(key string) bool {
	if _, ok := s.data[key]; ok {
		return true
	}

	return false
}

func (s *InMemoryStore[T]) logSetAction(action string, key string, value T) {
	s.rlw.LogSetAction(action, key, value)
}

func (s *InMemoryStore[T]) logDeleteAction(action string, key string) {
	s.rlw.LogDeleteAction(action, key)
}
