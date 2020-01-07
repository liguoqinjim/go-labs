package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"time"
)

func main() {
	//lab001()

	//lab002()

	lab003()
}

func lab001() {
	limiter := rate.NewLimiter(5, 1)

	for i := 0; i < 10; i++ {
		if err := limiter.Wait(context.TODO()); err != nil {
			log.Println(err)
		}

		log.Println(i)
	}
}

func lab002() {
	limiter := rate.NewLimiter(2, 1)

	for i := 0; i < 10; i++ {
		log.Println(limiter.Allow())
	}
}

func lab003() {
	limiter := rate.NewLimiter(2, 1)

	for i := 0; i < 10; i++ {
		go func(num int) {
			r := limiter.Reserve()

			if !r.OK() {
				log.Println("ok is false", num)
				return
			}

			log.Println(r.Delay())
			time.Sleep(r.Delay())
			log.Println("act", num)
		}(i)

		time.Sleep(time.Millisecond * 200)
	}

	time.Sleep(time.Hour)
}
