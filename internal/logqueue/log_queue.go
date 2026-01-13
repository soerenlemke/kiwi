package logqueue

import (
	"kiwi/internal/domain"
)

// LogQueue is a generic FIFO queue for log entries.
//
// Type parameter:
//   - T: A type that satisfies domain.AllowedTypes.
type LogQueue[T domain.AllowedTypes] struct {
	// Entries stores the queued log entries in FIFO order.
	Entries []domain.LogEntry[T]
}

// NewLogQueue creates and returns a new, empty LogQueue.
func NewLogQueue[T domain.AllowedTypes]() *LogQueue[T] {
	return &LogQueue[T]{
		Entries: make([]domain.LogEntry[T], 0),
	}
}

// Enqueue appends a new log entry to the end of the queue.
//
// Note: There is currently no queue size limit enforced.
func (l *LogQueue[T]) Enqueue(entry domain.LogEntry[T]) {
	l.Entries = append(l.Entries, entry)
}

// Dequeue removes and returns the first log entry in the queue.
//
// It returns nil if the queue is empty.
func (l *LogQueue[T]) Dequeue() *domain.LogEntry[T] {
	if l.isEmpty() {
		return nil
	}

	entry := l.Entries[0]
	l.Entries = l.Entries[1:]

	return &entry
}

// isEmpty reports whether the queue contains no entries.
func (l *LogQueue[T]) isEmpty() bool {
	return len(l.Entries) == 0
}

// size returns the current number of entries in the queue.
func (l *LogQueue[T]) size() int {
	return len(l.Entries)
}
