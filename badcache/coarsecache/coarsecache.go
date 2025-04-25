package coarsecache

import (
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	data  map[string]string
	mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Put(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
