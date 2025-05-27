package nonblocking

import (
	"sync/atomic"
)

type cacheSnapshot struct {
	data    map[string]any
	version uint64
}

type NonBlockingCache struct {
	state atomic.Value
}

func NewNonBlockingCache() *NonBlockingCache {
	c := &NonBlockingCache{}
	c.state.Store(&cacheSnapshot{
		data:    make(map[string]any),
		version: 0,
	})
	return c
}

func (c *NonBlockingCache) Set(key string, value any) {
	for {
		old := c.state.Load().(*cacheSnapshot)
		newData := make(map[string]any, len(old.data)+1)
		for k, v := range old.data {
			newData[k] = v
		}
		newData[key] = value

		newState := &cacheSnapshot{
			data:    newData,
			version: old.version + 1,
		}

		if c.state.CompareAndSwap(old, newState) {
			break
		}
	}
}

func (c *NonBlockingCache) Get(key string) (any, bool) {
	state := c.state.Load().(*cacheSnapshot)
	val, ok := state.data[key]
	return val, ok
}
