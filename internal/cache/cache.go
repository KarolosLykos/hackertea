// Package cache provides an interface for caching data.
package cache

import (
	"sync"

	"github.com/KarolosLykos/hackertea/internal/item"
)

// Cache is an interface for a cache.
type Cache interface {
	Get(key int) (*item.Item, bool)
	Set(key int, value *item.Item)
}

// MemCache is an implementation of the Cache interface that stores data in memory.
type MemCache struct {
	lock  sync.Mutex
	items map[int]*item.Item
}

// New returns a new MemCache.
func New() *MemCache {
	return &MemCache{items: make(map[int]*item.Item)}
}

// Get retrieves an item from the cache with the given key.
// Returns the item and a bool indicating whether the item was found.
func (m *MemCache) Get(key int) (*item.Item, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	v, ok := m.items[key]
	if !ok {
		return nil, false
	}

	return v, true
}

// Set sets an item in the cache with the given key and value.
func (m *MemCache) Set(key int, value *item.Item) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.items[key] = value
}
