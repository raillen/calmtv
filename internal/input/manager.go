package input

import "sync"

// Manager serializes semantic actions before dispatching them to subscribers.
type Manager struct {
	mu          sync.RWMutex
	subscribers []func(Event)
}

func NewManager() *Manager { return &Manager{} }

func (m *Manager) Subscribe(handler func(Event)) {
	if handler == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.subscribers = append(m.subscribers, handler)
}

func (m *Manager) Emit(event Event) {
	m.mu.RLock()
	subscribers := append([]func(Event){}, m.subscribers...)
	m.mu.RUnlock()

	for _, handler := range subscribers {
		handler(event)
	}
}
