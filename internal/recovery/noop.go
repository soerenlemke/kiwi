package recovery

type NopRlw[T any] struct{}

func (NopRlw[T]) LogSetAction(string, string, T) {}
func (NopRlw[T]) LogDeleteAction(string, string) {}
