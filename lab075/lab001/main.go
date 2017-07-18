package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go sayHello1()

	<-sigs
	log.Println("end")
}

func sayHello1() {
	i := 0
	for i < 10 {
		log.Println("i=", i)

		i++
		if i == 5 {
			go sayHello2()
		}

		if i == 7 {
			break
		}

		time.Sleep(time.Second * 2)
	}
}

func sayHello2() {
	j := 0
	for j < 10 {
		log.Println("sayHello2:", j)

		j++
		time.Sleep(time.Second * 1)
	}
}
