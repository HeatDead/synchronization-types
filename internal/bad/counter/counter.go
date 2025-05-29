package counter

import (
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		counter++ // Не защищенный доступ к переменной
	}
	wg.Done()
}
