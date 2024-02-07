package sync

import (
	"fmt"
	"sync"
)

func Run() {
	RunWaitGroup()
	RunManyFunc()
}

func RunWaitGroup() {
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

func RunManyFunc() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		fmt.Println("First")
	}()

	go func() {
		wg.Add(1)
		fmt.Println("Second")
	}()

	wg.Wait()
}
