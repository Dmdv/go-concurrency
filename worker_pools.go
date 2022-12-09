package main

import (
	"fmt"
	"time"
)

// https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8

func boris() {
	var workers = make(chan struct{}, 5)
	for i := 1; i <= 10; i++ {
		workers <- struct{}{}
		go func(job int) {
			defer func() {
				<-workers
			}()
			time.Sleep(1 * time.Second)
			fmt.Println(job)
		}(i)
	}
	//time.Sleep(2 * time.Second)
}
