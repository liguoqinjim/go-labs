package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			EchoNumber(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("end")
}

func EchoNumber(i int) {
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	log.Println("number=", i)
}
