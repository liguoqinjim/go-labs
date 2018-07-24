package main

import (
	"log"
	"time"
)

func worker(done chan bool) {
	log.Println("working...")
	time.Sleep(time.Second * 5)
	log.Println("done")

	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done
	log.Println("main done")
}
