package main

import (
	"log"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	log.Println("now_time=", time.Now())

	<-timer1.C
	log.Println("now_time=", time.Now())

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		log.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop() //you can cancel the timer before it expires
	if stop2 {
		log.Println("Timer 2 stopped")
	}
}
