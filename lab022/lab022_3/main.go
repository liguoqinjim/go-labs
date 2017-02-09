package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...", time.Now().Unix())
	time.Sleep(time.Second * 5)
	fmt.Println("done", time.Now().Unix())

	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done
	fmt.Println("main done", time.Now().Unix())
}
