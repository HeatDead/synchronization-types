package nonblocking

type NonBlockingCache struct {
	data map[string]any
}

func (c *NonBlockingCache) Set(key string, value any) {
	c.data[key] = value // Обычная операция, нет атомарности
}

func NewNonBlockingCache() *NonBlockingCache {
	c := &NonBlockingCache{}

	return c
}
