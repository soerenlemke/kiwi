package domain

type AllowedTypes interface {
	~string | ~bool | ~int | ~int64 | ~float64 | ~uint | ~uint64
}
