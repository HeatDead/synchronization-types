package badcache

import (
	"math/rand"
	"time"
)

// Неоптимальное решение

type Cache struct {
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	// Эмуляция работы с сетью/диском
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Put(key, value string) {
	// Эмуляция работы с сетью/диском
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	c.data[key] = value
}
