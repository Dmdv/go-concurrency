package main

import (
	"fmt"
	"github.com/dmdv/go-concurrency/channels"
	"github.com/dmdv/go-concurrency/maps"
	"github.com/dmdv/go-concurrency/sync"
)

func main() {
	// Channels

	fmt.Println("----- Channels examples ------	")
	channels.Run()

	// Worker pools

	fmt.Println("----- Worker pools examples ----")
	boris_pool()
	gammazero_pool()

	// Synchronisation

	fmt.Println("------- Synchronization examples --")
	sync.Run()

	// Maps

	fmt.Println("------- Map examples --")
	maps.Run()
	maps.Run_ReadWrite()
	maps.Run_Once()
}
