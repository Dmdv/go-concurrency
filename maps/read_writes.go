package maps

import (
	"fmt"
	"sync"
)

// Run_ReadWrite Use rwmutex for reads
func Run_ReadWrite() {
	storage := make(map[int]int, 500)

	wg := sync.WaitGroup{}
	reads := 1000
	writes := 1000
	mu := sync.RWMutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}
	wg.Add(reads)
	for i := 0; i < reads; i++ {
		i := 0
		go func() {
			defer wg.Done()

			mu.RLock()
			defer mu.RUnlock()
			_, _ = storage[i]
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
