package recovery

import (
	"kiwi/internal/domain"
)

type Recoverer interface {
	Recover()
}

type RecoverLogWriter[T domain.AllowedTypes] interface {
	LogSetAction(action string, key string, value T)
	LogDeleteAction(action string, key string)
}
