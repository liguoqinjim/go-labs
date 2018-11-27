package main

import (
	"log"
	"time"
)

func main() {
	time1 := time.Now()
	time.Sleep(time.Second * 10)
	time2 := time.Now()

	d1 := time.Since(time1)
	d2 := time.Since(time2)

	log.Println("d1=", d1)
	log.Println("d2=", d2)

	if d1 > time.Second*10 {
		log.Println("d1 > 10s")
	} else {
		log.Println("d1 <= 10s")
	}
}
