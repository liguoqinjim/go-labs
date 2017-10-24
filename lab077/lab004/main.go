package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	rw := new(sync.RWMutex)
	rw.RLock()
	log.Println("读锁1:Lock")

	//写锁1
	go func() {
		time.Sleep(time.Second)
		rw.Lock()
		log.Println("写锁1:Lock")

		time.Sleep(time.Second * 2)
		rw.Unlock()
		log.Println("写锁1:Unlock")
	}()

	time.Sleep(time.Second)
	rw.RUnlock()
	log.Println("读锁1:Unlock")

	time.Sleep(time.Second * 7)
	rw.RLock()
	log.Println("读锁1:再次Lock")

	go func() {
		rw.RLock()
		log.Println("读锁2:Lock")
		time.Sleep(time.Second * 10)
		rw.RUnlock()
		log.Println("读锁2:Unlock")
	}()

	time.Sleep(time.Second * 10)
	rw.RUnlock()
	log.Println("读锁1:Unlock")

	time.Sleep(time.Hour)
}
