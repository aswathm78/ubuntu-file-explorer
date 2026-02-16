package event

import (
	"sync"
)

// MemoryBus is a simple in-memory implementation of the Bus interface
type MemoryBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

// NewMemoryBus creates a new instance of MemoryBus
func NewMemoryBus() *MemoryBus {
	return &MemoryBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

// Publish sends an event to all subscribers of the given topic
func (b *MemoryBus) Publish(topic string, payload interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if chans, ok := b.subscribers[topic]; ok {
		for _, ch := range chans {
			// Non-blocking send to avoid stalling the publisher
			select {
			case ch <- payload:
			default:
				// If channel is full, drop the event (or log it)
				// For now, we drop to prevent blocking
			}
		}
	}
}

// Subscribe returns a channel that receives events for the given topic
func (b *MemoryBus) Subscribe(topic string) <-chan interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan interface{}, 100) // Buffer size of 100
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}
