package recovery

import (
	"kiwi/internal/domain"
	"kiwi/internal/log"
)

type Recoverer[T domain.AllowedTypes] interface {
	Recover()
	LogAction(action log.Action, key string, value T)
}
