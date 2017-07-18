package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Army struct {
	Id int
	X  int
	sync.Mutex
}

func (a *Army) Move() {
	a.X += 10
	log.Printf("army[%d] move to [%d]\n", a.Id, a.X)
}

func (a *Army) Scout() {
	log.Printf("army[%d] scout\n", a.Id)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	army := &Army{Id: 10001, X: 0}

	go func() {
		i := 0
		for i < 5 {
			army.Lock()
			army.Move()
			i++
			if i == 1 {
				go func() {
					log.Println("goroutine2:")
					j := 0
					for j < 3 {
						army.Lock()
						army.Scout()
						j++
						army.Unlock()
						time.Sleep(time.Second * 1)
					}
				}()
			}

			time.Sleep(time.Second * 9)
			army.Unlock() //这里不执行unlock操作的话会出错，之后的for会锁死
			time.Sleep(time.Second * 3)
		}
		log.Println("go end")
	}()

	<-sigs
	log.Println("end")
}
