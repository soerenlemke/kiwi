package recovery

// NopRlw is a no-op recovery logger currently used for testing.
type NopRlw[T any] struct{}

func (NopRlw[T]) LogSetAction(string, string, T) {}
func (NopRlw[T]) LogDeleteAction(string, string) {}
