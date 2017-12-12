package main

import (
	"expvar"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

var visits = expvar.NewInt("visits")

func main() {
	gr := expvar.NewInt("Goroutines")
	go func() {
		for range time.Tick(time.Second * 3) {
			gr.Set(int64(runtime.NumGoroutine()))
		}
	}()

	//创建goroutine
	for i := 0; i < 100; i++ {
		wg.Add(1)
		a := make(map[int]int)
		go func(a map[int]int) {
			for i := 0; i < 1000000; i++ {
				a[i] = 1
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}
			wg.Done()
		}(a)
	}

	http.ListenAndServe(":1234", nil)

	wg.Wait()

	log.Println("end")
}
