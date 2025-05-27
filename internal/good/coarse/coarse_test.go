package coarse

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

// Тесты
func hasMutex(c interface{}) bool {
	return reflect.ValueOf(c).Elem().FieldByName("mu").IsValid()
}

func TestCoarseLock(t *testing.T) {
	c := NewCoarseCache()

	// Проверка наличия мьютекса в структуре через рефлексию
	if !hasMutex(c) {
		t.Fatal("Coarse lock not implemented")
	}

	// Параллельные операции
	for i := 0; i < 1000; i++ {
		go c.Set("key", i)
		go c.Get("key")
	}
}

func TestCoarseCache(t *testing.T) {
	c := NewCoarseCache()

	if !hasMutex(c) {
		t.Fatal("Coarse lock not implemented")
	}

	// Конкурентные операции
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			c.Set(fmt.Sprintf("key%d", i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
