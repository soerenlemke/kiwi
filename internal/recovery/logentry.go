package recovery

import (
	"kiwi/internal/domain"
	"time"
)

type LogEntry[T domain.AllowedTypes] struct {
	TimeStamp time.Time
	key       string
	value     T
}

func NewLogEntry[T domain.AllowedTypes](key string, value T) *LogEntry[T] {
	return &LogEntry[T]{
		TimeStamp: time.Now(),
		key:       key,
		value:     value,
	}
}
