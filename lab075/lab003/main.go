package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Army struct {
	Aid       int
	Apos      int
	ReturnSig chan int
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	a := &Army{Aid: 10001, Apos: 0, ReturnSig: make(chan int, 1)}
	go a.Move()

	time.Sleep(time.Second * 3)
	a.ReturnSig <- 1

	<-sigs
	log.Println("end")
}

func (a *Army) Move() {
	log.Printf("Army[%d] start move\n", a.Aid)
	for {
		log.Println("Move的for循环")
		select {
		case <-a.ReturnSig:
			log.Printf("Army[%d] Return!\n", a.Aid)
		default:
			a.Apos += 10
			log.Printf("Army[%d].Apos=%d\n", a.Aid, a.Apos)
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}
	log.Printf("Army[%d] Move() quit\n", a.Aid)
}

func (a *Army) Stop() {
	log.Printf("Army[%d] stop!\n", a.Aid)
}
