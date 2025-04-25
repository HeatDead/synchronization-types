package nonblockcache

import (
	"math/rand"
	"sync/atomic"
	"time"
	"unsafe"
)

type entry struct {
	value string
}

type Cache struct {
	data unsafe.Pointer // *map[string]*entry
}

func NewCache() *Cache {
	m := make(map[string]*entry)
	return &Cache{
		data: unsafe.Pointer(&m),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	data := *(*map[string]*entry)(atomic.LoadPointer(&c.data))
	e, ok := data[key]
	if !ok {
		return "", false
	}
	return e.value, true
}

func (c *Cache) Put(key, value string) {
	for {
		oldData := *(*map[string]*entry)(atomic.LoadPointer(&c.data))
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		newData := make(map[string]*entry, len(oldData)+1)
		for k, v := range oldData {
			newData[k] = v
		}
		newData[key] = &entry{value: value}

		if atomic.CompareAndSwapPointer(
			&c.data,
			unsafe.Pointer(&oldData),
			unsafe.Pointer(&newData),
		) {
			return
		}
	}
}
