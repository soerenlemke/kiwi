package log

import (
	"kiwi/internal/domain"
	"time"
)

type Entry[T domain.AllowedTypes] struct {
	Timestamp time.Time
	Action    Action
	Key       string
	Value     T
}

func NewEntry[T domain.AllowedTypes](action Action, key string, value T) Entry[T] {
	return Entry[T]{
		Action:    action,
		Timestamp: time.Now(),
		Key:       key,
		Value:     value,
	}
}
