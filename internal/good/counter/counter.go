package counter

import (
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		mu.Lock() // Защита доступа к переменной
		counter++
		mu.Unlock()
	}
	wg.Done()
}
