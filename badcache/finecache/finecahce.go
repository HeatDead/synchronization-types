package finecache

import (
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	rwmu sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.rwmu.RLock()
	defer c.rwmu.RUnlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Put(key, value string) {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
