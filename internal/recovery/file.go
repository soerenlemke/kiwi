package recovery

import (
	"context"
	"kiwi/internal/domain"
	"kiwi/internal/logqueue"
	"os"
	"sync"
	"time"
)

type FileRecovery[T domain.AllowedTypes] struct {
	filePath string
	logQueue logqueue.LogQueue[T]
	file     *os.File
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

func New[T domain.AllowedTypes](filePath string) (*FileRecovery[T], error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	logQueue := logqueue.New[T]()

	fileRecovery := &FileRecovery[T]{
		filePath: filePath,
		logQueue: *logQueue,
		file:     file,
		cancel:   cancel,
	}

	fileRecovery.wg.Go(func() {
		fileRecovery.processQueue(ctx)
	})

	return fileRecovery, nil
}

func (f *FileRecovery[T]) LogSetAction(action string, key string, value T) {
	logEntry := domain.NewLogEntry[T](key, value)
	f.logQueue.Enqueue(logEntry)
}

func (f *FileRecovery[T]) LogDeleteAction(action string, key string) {
	//TODO implement me
	panic("implement me")
}

func (f *FileRecovery[T]) ReadLog() {
	//TODO implement me
	panic("implement me")
}

func (f *FileRecovery[T]) Recover() {
	//TODO implement me
	panic("implement me")
}

func (f *FileRecovery[T]) Close() error {
	f.cancel()
	f.wg.Wait()
	return f.file.Close()
}

func (f *FileRecovery[T]) processQueue(ctx context.Context) {
	defer f.wg.Done()
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			f.flushQueue()
			return
		case <-ticker.C:
			f.flushQueue()
		}
	}
}

func (f *FileRecovery[T]) flushQueue() {
	for !f.logQueue.IsEmpty() {
		entry := f.logQueue.Dequeue()
		if entry != nil {
			line := formatLogEntry(entry)
			_, _ = f.file.WriteString(line + "\n")
		}
	}
}

func formatLogEntry[T domain.AllowedTypes](entry *domain.LogEntry[T]) string {
	// TODO: JSON oder eigenes Format implementieren
	return entry.TimeStamp.Format(time.RFC3339) + " key=... value=..."
}

/*
func formatLogEntry[T domain.AllowedTypes](entry *domain.LogEntry[T]) string {
    // TODO: JSON oder eigenes Format implementieren
    return entry.TimeStamp.Format(time.RFC3339) + " key=... value=..."
}

func (f *FileRecovery[T]) Close() error {
    f.cancel()
    f.wg.Wait()
    return f.file.Close()
}
*/
