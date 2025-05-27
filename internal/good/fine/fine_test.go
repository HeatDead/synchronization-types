package fine

import "testing"

func TestFineGrained(t *testing.T) {
	c := NewFineCache()

	// Должно быть 16 шардов
	if len(c.shards) != 16 {
		t.Error("Fine-grained sharding incorrect")
	}

	// Параллельная запись в разные шарды
	keys := []string{"a", "b", "c", "d"}
	for _, key := range keys {
		go c.Set(key, 1)
	}
}
