package recovery

import (
	"kiwi/internal/domain"
)

type Recoverer[T domain.AllowedTypes] interface {
	Recover()
	ReadLog()
	LogSetAction(action string, key string, value T)
	LogDeleteAction(action string, key string)
}
