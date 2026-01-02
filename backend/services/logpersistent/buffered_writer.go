package logpersistent

import (
	"sync"
	"time"
)

const (
	bufferedMaxCount = 100
	bufferedInterval = 5 * time.Second
)

type bufferedSender func(rows [][]byte)

type bufferedWriter struct {
	buffer   [][]byte
	maxCount int
	mu       sync.Mutex
	ticker   *time.Ticker
	done     chan struct{}
	send     bufferedSender
}

func newBufferedWriter(send bufferedSender) *bufferedWriter {
	return &bufferedWriter{
		buffer:   make([][]byte, 0),
		maxCount: bufferedMaxCount,
		ticker:   time.NewTicker(bufferedInterval),
		done:     make(chan struct{}),
		send:     send,
	}
}

func (w *bufferedWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	entry := make([]byte, len(p))
	copy(entry, p)

	w.buffer = append(w.buffer, entry)

	if len(w.buffer) >= w.maxCount {
		w.flush()
	}

	return len(p), nil
}

func (w *bufferedWriter) flushLoop() {
	for {
		select {
		case <-w.ticker.C:
			w.mu.Lock()
			w.flush()
			w.mu.Unlock()
		case <-w.done:
			return
		}
	}
}

func (w *bufferedWriter) flush() {
	if len(w.buffer) == 0 {
		return
	}

	w.send(w.buffer)
	w.buffer = w.buffer[:0]
}

func (w *bufferedWriter) Close() error {
	close(w.done)
	w.ticker.Stop()

	w.mu.Lock()
	w.flush()
	w.mu.Unlock()

	return nil
}
