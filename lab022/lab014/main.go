package main

import (
	"log"
	"time"
)

func main() {
	jobs := make(chan int, 5)

	go func() {
		for {
			job, ok := <-jobs
			log.Println("job receive", job)
			if !ok {
				log.Println("jobs close receive")
				break
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		log.Println("sent job", j)
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 3)
	log.Println("close jobs")
	close(jobs)

	time.Sleep(time.Hour)
}
