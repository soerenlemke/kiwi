package store

type AllowedTypes interface {
	~string | ~bool | ~int | ~int64 | ~float64 | ~uint | ~uint64
}

type Store[T AllowedTypes] interface {
	Set(key string, value T) error
	Get(key string) (T, bool, error) // bool = key exists
	Delete(key string) error
}
