package optimistic

import "sync"

type OptimisticCache struct {
	data map[string]any
	mu   sync.Mutex
}

func (c *OptimisticCache) Update(key string, newValue any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = newValue // Нет проверки версий!
	return true
}

func NewOptimisticCache() *OptimisticCache {
	return &OptimisticCache{}
}
func (c *OptimisticCache) Set(key string, value any) {

}

func (c *OptimisticCache) Get(key string) (any, uint64) {
	return 0, 0
}
