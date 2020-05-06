package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	m := make(map[int]string)
	for i := 0; i < 1000; i++ {
		m[i] = strconv.Itoa(i)
	}

	for i := 0; i < 10000; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				a := m[i]

				if i%100 == 0 {
					fmt.Print(a)
				}
			}
		}()
	}

	//写会报错
	//for i := 0;i < 100;i++{
	//	go func() {
	//		for i := 0;i < 20;i++{
	//			m[i] = "1"
	//		}
	//	}()
	//}

	//delete会报错
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				delete(m, i)
				time.Sleep(time.Millisecond * 200)
			}
		}()
	}

	<-sigs
}
