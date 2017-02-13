package main

import (
	"fmt"
	"time"
)

func main() {
	//normal rate limiting
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter //limiting
		fmt.Println("request", req, time.Now())
	}

	//2 burstyLimiter bustryLimiter事先就<-了三个，那么前三个就不用等候这个rate limiting了
	fmt.Println()
	bustryLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		bustryLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			fmt.Println("li ", time.Now())
			bustryLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyRequests
		fmt.Println("request", req, time.Now())
	}
}
