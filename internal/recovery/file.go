package recovery

import "kiwi/internal/domain"

type FileRecovery[T domain.AllowedTypes] struct {
	filePath string
}

func NewFileRecovery[T domain.AllowedTypes](filePath string) *FileRecovery[T] {
	return &FileRecovery[T]{
		filePath: filePath,
	}
}

func (f FileRecovery[T]) LogSetAction(action string, key string, value T) {
	//TODO implement me
	panic("implement me")
}

func (f FileRecovery[T]) LogDeleteAction(action string, key string) {
	//TODO implement me
	panic("implement me")
}
