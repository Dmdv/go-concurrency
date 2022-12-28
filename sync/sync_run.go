package sync

import (
	"fmt"
	"sync"
)

func Run() {
	counter := 20
	wg := &sync.WaitGroup{}
	for i := 0; i < counter; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i * i)
		}()
	}

	wg.Wait()
}
