package lazy

type LazyCache struct {
	data map[string]any
}

func (c *LazyCache) Set(key string, value any) {
	if c.data == nil { // Гонка при инициализации
		c.data = make(map[string]any)
	}
	c.data[key] = value
}
