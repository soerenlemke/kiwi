package recovery

// NopRlw is a no-op recovery logger currently used for testing.
type NopRlw[T any] struct{}

func (t NopRlw[T]) Recover()                                        {}
func (t NopRlw[T]) ReadLog()                                        {}
func (t NopRlw[T]) LogSetAction(action string, key string, value T) {}
func (NopRlw[T]) LogDeleteAction(string, string)                    {}
