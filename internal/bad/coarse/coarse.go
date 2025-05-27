package coarse

type CoarseCache struct {
	data map[string]any
}

func NewCoarseCache() *CoarseCache {
	return &CoarseCache{
		data: make(map[string]any),
	}
}

// Нет мьютекса!
func (c *CoarseCache) Set(key string, value any) {
	c.data[key] = value // ГОНКА ДАННЫХ
}

func (c *CoarseCache) Get(key string) (any, bool) {
	val, ok := c.data[key]
	return val, ok
}
