package fine

import "sync"

type FineCache struct {
	shards map[string]map[string]any
	mu     sync.Mutex // Только один мьютекс для всех шардов
}

func NewFineCache() *FineCache {
	return &FineCache{
		shards: make(map[string]map[string]any),
	}
}

func (c *FineCache) Set(key string, value any) {
	c.mu.Lock() // Неправильно: блокирует все шарды
	defer c.mu.Unlock()
	shardKey := key[:1]
	if c.shards[shardKey] == nil {
		c.shards[shardKey] = make(map[string]any)
	}
	c.shards[shardKey][key] = value
}
