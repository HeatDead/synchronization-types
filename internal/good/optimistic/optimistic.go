package optimistic

import (
	"sync"
	"sync/atomic"
	"time"
)

type OptimisticCache struct {
	mu   sync.RWMutex
	data map[string]*entry
}

type entry struct {
	value   any
	version uint64
}

func NewOptimisticCache() *OptimisticCache {
	return &OptimisticCache{
		data: make(map[string]*entry),
	}
}

func (c *OptimisticCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = &entry{value: value, version: 0}
}

func (c *OptimisticCache) Get(key string) (any, uint64) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if e, ok := c.data[key]; ok {
		return e.value, atomic.LoadUint64(&e.version)
	}
	return nil, 0
}

func (c *OptimisticCache) Update(key string, newValue any) bool {
	c.mu.RLock()
	e, ok := c.data[key]
	if !ok {
		c.mu.RUnlock()
		return false
	}
	oldVersion := atomic.LoadUint64(&e.version)
	c.mu.RUnlock()

	// Имитация долгой операции
	time.Sleep(time.Millisecond)

	c.mu.Lock()
	defer c.mu.Unlock()

	if e := c.data[key]; e == nil || e.version != oldVersion {
		return false
	}

	c.data[key].value = newValue
	atomic.StoreUint64(&c.data[key].version, oldVersion+1)
	return true
}
