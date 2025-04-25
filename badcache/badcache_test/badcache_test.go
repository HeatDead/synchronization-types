package badcache_test

import (
	"strconv"
	"sync"
	"synchronization-types/badcache"
	"testing"
)

func TestCacheRaceCondition(t *testing.T) {
	cache := badcache.NewCache()
	var wg sync.WaitGroup

	// Запускаем несколько горутин для записи
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Put("key", "value"+strconv.Itoa(i))
		}(i)
	}

	// Запускаем несколько горутин для чтения
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Get("key")
		}()
	}

	wg.Wait()

	// Проверяем, что в кэше осталось какое-то значение
	val, ok := cache.Get("key")
	if !ok {
		t.Error("Key not found after concurrent access")
	}
	t.Logf("Final value: %s", val)
}

func TestCacheConsistency(t *testing.T) {
	cache := badcache.NewCache()
	const workers = 50
	const iterations = 100
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := "key" + strconv.Itoa(j%10)
				cache.Put(key, "value"+strconv.Itoa(j))
				if _, ok := cache.Get(key); !ok {
					t.Errorf("Key %s lost", key)
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestParallelOperations(t *testing.T) {
	cache := badcache.NewCache()
	const keys = 10
	const workers = 20

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := "key" + strconv.Itoa(j%keys)
				if workerID%2 == 0 {
					cache.Put(key, "value"+strconv.Itoa(workerID)+"-"+strconv.Itoa(j))
				} else {
					cache.Get(key)
				}
			}
		}(i)
	}

	wg.Wait()
}
