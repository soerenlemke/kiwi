package store

import (
	"kiwi/internal/domain"
)

type Store[T domain.AllowedTypes] interface {
	Set(key string, value T) error
	Get(key string) (T, bool, error) // bool = key exists
	Delete(key string) error
	logSetAction(action string, key string, value T)
	logDeleteAction(action string, key string)
}
