package lazy

import "sync"

type LazyCache struct {
	initOnce sync.Once
	data     map[string]any
	mu       sync.RWMutex
}

func (c *LazyCache) init() {
	c.data = make(map[string]any)
}

func (c *LazyCache) Set(key string, value any) {
	c.initOnce.Do(c.init)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
