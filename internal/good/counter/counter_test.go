package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var wg sync.WaitGroup
	counter = 0 // Сброс счетчика перед тестом

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	if counter != 10000 { // Ожидаемое значение
		t.Errorf("Expected counter to be 10000, got %d", counter)
	}
}
