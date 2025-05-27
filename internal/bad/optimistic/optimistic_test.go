package optimistic

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestOptimisticLock(t *testing.T) {
	c := NewOptimisticCache()
	c.Set("key", "initial")

	var success int32 // Используем int32 для атомарных операций
	var wg sync.WaitGroup
	start := make(chan struct{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			<-start
			if c.Update("key", "new") {
				atomic.AddInt32(&success, 1) // Корректная атомарная операция
			}
			wg.Done()
		}()
	}

	close(start)
	wg.Wait()

	if atomic.LoadInt32(&success) != 1 { // Атомарное чтение
		t.Errorf("Expected 1 success, got %d", atomic.LoadInt32(&success))
	}
}
