package main

import (
	"log"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//log.Println("worker", id, "started job", j, "now_time", time.Now())
		log.Printf("worker[%d] started job[%d] nowtime[%v]", id, j, time.Now())
		time.Sleep(time.Second)
		log.Println("worker", id, "finished job", j, "now_time", time.Now())
		log.Printf("workder[%d] finished job[%d] nowtime[%v]", id, j, time.Now())
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		//<-results
		log.Println("result", <-results)
	}
}
