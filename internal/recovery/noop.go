package recovery

import (
	"kiwi/internal/log"
)

// NopRlw is a no-op recovery logger currently used for testing.
type NopRlw[T any] struct{}

func (t NopRlw[T]) Recover()                                         {}
func (t NopRlw[T]) LogAction(action log.Action, key string, value T) {}
