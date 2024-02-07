package timers

import (
	"fmt"
	"log"
	"time"
)

func Run() {
	RunTimeAfter()
	RunTimeOut()
	RepeatEverySecond()
	Foo()
}

func Foo() {
	timer := time.AfterFunc(time.Minute, func() {
		log.Println("Foo run for more than a minute.")
	})
	defer timer.Stop()
	// Do heavy work
}

func RepeatEverySecond() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fmt.Println("Tick")
		}
	}()
}

func RunTimeAfter() {
	AFP := make(chan string)
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(time.Hour):
		fmt.Println("No news in an hour.")
	}
}

func RunTimeOut() {
	AFP := make(chan string)
	for alive := true; alive; {
		timer := time.NewTimer(time.Hour)
		select {
		case news := <-AFP:
			timer.Stop()
			fmt.Println(news)
		case <-timer.C:
			alive = false
			fmt.Println("No news in an hour. Service aborting.")
		}
	}
}
