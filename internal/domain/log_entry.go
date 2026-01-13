package domain

import (
	"time"
)

type LogEntry[T AllowedTypes] struct {
	TimeStamp time.Time
	key       string
	value     T
}

func NewLogEntry[T AllowedTypes](key string, value T) LogEntry[T] {
	return LogEntry[T]{
		TimeStamp: time.Now(),
		key:       key,
		value:     value,
	}
}
