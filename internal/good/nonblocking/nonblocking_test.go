package nonblocking

import (
	"synchronization-types/internal/test_utils"
	"testing"
)

func TestNonBlocking(t *testing.T) {
	c := NewNonBlockingCache()

	// Проверка на использование atomic.Value
	if test_utils.HasMutex(c) {
		t.Error("Mutex used in non-blocking impl")
	}

	// Параллельные обновления
	for i := 0; i < 100; i++ {
		go c.Set("key", i)
	}
}
