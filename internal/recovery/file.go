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
	logEntry := NewLogEntry[T](key, value)
	println(logEntry)

	// TODO: write to logging queue -> queue writes to file with mutex etc.
}

func (f FileRecovery[T]) LogDeleteAction(action string, key string) {
	//TODO implement me
	panic("implement me")
}

func (f FileRecovery[T]) ReadLog() {
	//TODO implement me
	panic("implement me")
}

func (f FileRecovery[T]) Recover() {
	//TODO implement me
	panic("implement me")
}
