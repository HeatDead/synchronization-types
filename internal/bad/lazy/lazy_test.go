package lazy

import (
	"reflect"
	"sync"
	"testing"
)

func TestLazyRace(t *testing.T) {
	c := &LazyCache{}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Set("key", "value")
		}()
	}
	wg.Wait()

	// Проверка на использование sync.Once
	if !hasSyncOnce(c) {
		t.Error("sync.Once not used")
	}
}

// Вспомогательная функция для проверки
func hasSyncOnce(c interface{}) bool {
	return reflect.ValueOf(c).Elem().FieldByName("initOnce").IsValid()
}
