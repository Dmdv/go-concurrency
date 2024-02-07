package main

import (
	"fmt"
	"time"
)

// https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8

// boris_pool is a function that demonstrates how to use a pool of workers
func boris_pool() {
	var workers = make(chan struct{}, 5)
	for i := 1; i <= 10; i++ {
		workers <- struct{}{}

		// This is a job unit. It blocks if the workers channel is full
		go func(job int) {
			defer func() {
				// Release the worker
				<-workers
			}()
			time.Sleep(time.Duration(job) * time.Second)
			fmt.Printf("Boris pool Job# %v done\n", job)
		}(i)
	}
}
