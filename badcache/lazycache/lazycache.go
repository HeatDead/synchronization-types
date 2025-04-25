package lazycache

import (
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	data   map[string]string
	mu     sync.Mutex
	inited bool
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) init() {
	if !c.inited {
		c.mu.Lock()
		defer c.mu.Unlock()
		if !c.inited {
			c.data = make(map[string]string)
			c.inited = true
		}
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.init()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Put(key, value string) {
	c.init()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
