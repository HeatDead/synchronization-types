package optimisticcache

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type entry struct {
	value   string
	version uint64
}

type Cache struct {
	data    map[string]*entry
	version uint64
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]*entry),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	e, ok := c.data[key]
	if !ok {
		return "", false
	}
	return e.value, true
}

func (c *Cache) Put(key, value string) {
	for {
		oldVersion := atomic.LoadUint64(&c.version)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		e := c.data[key]
		if e == nil {
			e = &entry{value: value, version: oldVersion}
			c.data[key] = e
			return
		}

		if atomic.CompareAndSwapUint64(&e.version, oldVersion, oldVersion+1) {
			e.value = value
			atomic.AddUint64(&c.version, 1)
			return
		}
	}
}
