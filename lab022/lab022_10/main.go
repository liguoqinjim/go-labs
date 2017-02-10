package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("now_time=", time.Now().Unix())

	<-timer1.C
	fmt.Println("now_time=", time.Now().Unix())

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop() //you can cancel the timer before it expires
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
