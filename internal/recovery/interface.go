package recovery

import (
	"kiwi/internal/domain"
)

type Recoverer[T domain.AllowedTypes] interface {
	Recover()
	RecoverLogReader[T]
	RecoverLogWriter[T]
}

type RecoverLogReader[T domain.AllowedTypes] interface {
	ReadLog()
}

type RecoverLogWriter[T domain.AllowedTypes] interface {
	LogSetAction(action string, key string, value T)
	LogDeleteAction(action string, key string)
}
