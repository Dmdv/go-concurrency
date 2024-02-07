package channels

import (
	"fmt"
	"time"
)

func Run() {
	RunTestClosedChannel()
	RunBlocker1()
	RunBlocker2()
	<-PublishAsync()
	RunIntSequence()
	RunTimout()
	RunKillRoutine()
}

// 1. Channel used to block the main goroutine until all the workers are done

func RunBlocker1() {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker 1 completed")
	}()

	<-ch
}

func RunBlocker2() {
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		defer close(ch)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker 2 completed")
	}(ch)

	<-ch
}

func PublishAsync() chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		time.Sleep(1 * time.Second)
		fmt.Println("PublishAsync completed")
	}()

	return ch
}

func RunTestClosedChannel() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
		fmt.Println("Closed channel")
	}()

	fmt.Println(<-ch) // Print "Hello!".
	fmt.Println(<-ch) // Print the zero value "" without blocking.
	fmt.Println(<-ch) // Once again print "".
	fmt.Println(<-ch) // Once again print "".
	v, ok := <-ch     // v is "", ok is false.
	_ = v
	_ = ok

	// Receive values from ch until closed.
	for v := range ch {
		fmt.Println(v) // Will not be executed.
	}
}

func RunTimout() {
	AFP := make(chan string)
	defer close(AFP)
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(time.Second):
		fmt.Println("Time out: No news in one second")
	}
}

// Kill goroutine with a timeout

func RunKillRoutine() {
	ch := NumGenerator()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 0
	//fmt.Println(<-ch)
}

func NumGenerator() chan int {
	ch := make(chan int)
	go func() {
		n := 0
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

// 2. Random binary sequence generator

func RunIntSequence() {
	PrintRandomIntSequence(0)
	PrintRandomIntSequence(1)
	PrintRandomIntSequence(1)
	PrintRandomIntSequence(1)

	for i := 0; i < 10; i++ {
		PrintRandomIntSequence(2)
	}

	for i := 0; i < 10; i++ {
		PrintRandomIntSequence(8)
	}
}

func PrintRandomIntSequence(len int) {
	ch := RandomBinarySequence(len)
	val := BitArrayToInt(ch)
	fmt.Printf("From len = %v, Value = %v\n", len, val)
}

func RandomBinarySequence(len int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < len; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}

func BitArrayToInt(ch chan int) int {
	n := 0
	for i := range ch {
		n = n<<1 | i
	}
	return n
}
