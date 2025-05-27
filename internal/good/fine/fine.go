package fine

import "sync"

const shardCount = 16

type FineCache struct {
	shards []*shard
}

type shard struct {
	data map[string]any
	mu   sync.RWMutex
}

func NewFineCache() *FineCache {
	shards := make([]*shard, shardCount)
	for i := range shards {
		shards[i] = &shard{data: make(map[string]any)}
	}
	return &FineCache{shards: shards}
}

func (c *FineCache) getShard(key string) *shard {
	return c.shards[fnv32(key)%shardCount]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for _, c := range key {
		hash *= 16777619
		hash ^= uint32(c)
	}
	return hash
}

func (c *FineCache) Set(key string, value any) {
	shard := c.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.data[key] = value
}
