package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func fakeReadDB(i int) {
	go func() {
		ms := make([]int, i)
		for n, v := range ms {
			ms[n] = v + 1
		}
		time.Sleep(time.Minute * 2)
		log.Println(ms)
	}()
}

func main() {
	//监听pprof
	go func() {
		http.ListenAndServe("localhost:7777", nil)
	}()

	//操作
	for i := 0; i < 100; i++ {
		go fakeReadDB(100)
		time.Sleep(time.Second * 3)
	}

	time.Sleep(time.Hour)
}
