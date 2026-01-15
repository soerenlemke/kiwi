package recovery

import (
	"context"
	"kiwi/internal/domain"
	"kiwi/internal/queue"
	"log/slog"
	"os"
	"sync"
	"time"
)

const flushInterval = 10 * time.Millisecond

type FileRecovery[T domain.AllowedTypes] struct {
	logger   *slog.Logger
	filePath string
	logQueue *queue.LogQueue[T]
	file     *os.File
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

func NewFileRecovery[T domain.AllowedTypes](logger *slog.Logger, filePath string) (*FileRecovery[T], error) {
	if logger == nil {
		logger = slog.Default()
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	logQueue := queue.NewLogQueue[T](logger)

	fileRecovery := &FileRecovery[T]{
		logger:   logger,
		filePath: filePath,
		logQueue: logQueue,
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
	ticker := time.NewTicker(flushInterval)
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
