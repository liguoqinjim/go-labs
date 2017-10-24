package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	rw := new(sync.RWMutex)
	var deadLockCase time.Duration = 1
	go func() {
		time.Sleep(time.Second * deadLockCase)
		log.Println("Writer Try")
		rw.Lock()
		log.Println("Writer Fetch")
		time.Sleep(time.Second * 1)
		log.Println("Writer Release")
		rw.Unlock()
	}()
	log.Println("Reader 1 Try")
	rw.RLock()
	log.Println("Reader 1 Fetch")
	time.Sleep(time.Second * 2)
	log.Println("Reader 2 Try")
	rw.RLock()
	log.Println("Reader 2 Fetch")
	time.Sleep(time.Second * 2)
	log.Println("Reader 1 Release")
	rw.RUnlock()
	time.Sleep(time.Second * 1)
	log.Println("Reader 2 Release")
	rw.RUnlock()
	time.Sleep(time.Second * 2)
	log.Println("Done")
}
