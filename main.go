package main

import (
	"fmt"
	"github.com/dmdv/go-concurrency/maps"
	"github.com/dmdv/go-concurrency/sync"
)

func main() {
	fmt.Println("Worker pools examples")
	boris_pool()
	gammazero_pool()

	fmt.Println("Synchronization examples")

	sync.Run()

	maps.Run()
	maps.Run_ReadWrite()
	maps.Run_Once()
}
