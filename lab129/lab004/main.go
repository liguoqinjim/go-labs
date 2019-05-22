package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Class struct {
	Students map[string]int
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	var classes sync.Map

	for i := 0; i < 10; i++ {
		class := &Class{Students: make(map[string]int)}

		for j := 0; j < 20; j++ {
			class.Students[fmt.Sprintf("student_%d", j)] = j
		}

		classes.Store(i, class)
	}

	go func() {
		if class, ok := classes.Load(0); !ok {
			log.Fatalf("classes.Load not exist")
		} else {
			log.Println("goroutine1 lock")
			time.Sleep(time.Second * 15)

			log.Println("goroutine1 unlock")
			_ = class
		}
	}()

	time.Sleep(time.Second)

	go func() {
		if class, ok := classes.Load(0); !ok {
			log.Fatalf("classes.Load not exist")
		} else {
			log.Println("goroutine2 lock")
			time.Sleep(time.Second * 15)

			log.Println("goroutine2 unlock")
			_ = class
		}
	}()

	//!!!
	//这时候在goroutine1和goroutine2里面同时操作key为0的那个class的话，还是会报错的

	<-sigs
	log.Println("end")
}
