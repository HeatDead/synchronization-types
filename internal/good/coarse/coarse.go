package coarse

import (
	"sync"
)

type CoarseCache struct {
	data map[string]any
	mu   sync.RWMutex
}

func NewCoarseCache() *CoarseCache {
	return &CoarseCache{
		data: make(map[string]any),
	}
}

func (c *CoarseCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *CoarseCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}
