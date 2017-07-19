package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	//声明
	var mutex sync.Mutex
	log.Println("Lock the lock. (G0)")
	//加锁mutex
	mutex.Lock()

	log.Println("The lock is locked.(G0)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			log.Printf("Lock the lock. (G%d)\n", i)
			mutex.Lock()
			log.Printf("The lock is locked. (G%d)\n", i)
		}(i)
	}
	//休息一会,等待打印结果
	time.Sleep(time.Second)
	log.Println("Unlock the lock. (G0)")
	//解锁mutex
	mutex.Unlock()

	log.Println("The lock is unlocked. (G0)")
	//休息一会,等待打印结果
	time.Sleep(time.Second * 3)
}
