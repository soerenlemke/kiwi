package queue

import (
	"kiwi/internal/domain"
	"kiwi/internal/log"
	"log/slog"
)

// LogQueue is a generic FIFO queue for log entries.
//
// Type parameter:
//   - T: A type that satisfies domain.AllowedTypes.
type LogQueue[T domain.AllowedTypes] struct {
	logger *slog.Logger
	// Entries stores the queued log entries in FIFO order.
	Entries []log.Entry[T]
}

// NewLogQueue creates and returns a new, empty LogQueue.
func NewLogQueue[T domain.AllowedTypes](logger *slog.Logger) *LogQueue[T] {
	if logger == nil {
		logger = slog.Default()
	}

	return &LogQueue[T]{
		logger:  logger,
		Entries: make([]log.Entry[T], 0),
	}
}

// Enqueue appends a new log entry to the end of the queue.
//
// Note: There is currently no queue size limit enforced.
func (l *LogQueue[T]) Enqueue(entry log.Entry[T]) {
	l.Entries = append(l.Entries, entry)
	l.logger.Debug("Enqueued", "entry", entry)
}

// Dequeue removes and returns the first log entry in the queue.
//
// It returns nil if the queue is empty.
func (l *LogQueue[T]) Dequeue() *log.Entry[T] {
	if l.IsEmpty() {
		l.logger.Debug("Queue is empty")

		return nil
	}

	entry := l.Entries[0]
	l.Entries = l.Entries[1:]

	l.logger.Debug("Dequeued", "entry", entry)

	return &entry
}

// IsEmpty checks if the LogQueue is empty.
//
// Returns:
//   - true: If the queue has no entries.
//   - false: If the queue contains one or more entries.
func (l *LogQueue[T]) IsEmpty() bool {
	return len(l.Entries) == 0
}
