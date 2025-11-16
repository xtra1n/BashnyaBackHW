package concurrent_map

import "sync"

type SafeMap struct {
	data map[string]int
	mu   sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (m *SafeMap) Set(key string, value int) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.data[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
	m.mu.RLock()
	value, isSearched := m.data[key]
	m.mu.RUnlock()

	return value, isSearched
}

func (m *SafeMap) Len() int {
	m.mu.RLock()
	len := len(m.data)
	m.mu.RUnlock()

	return len
}
