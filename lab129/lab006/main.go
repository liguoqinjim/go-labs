package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	m := make(map[int]string)
	for i := 0; i < 20; i++ {
		m[i] = strconv.Itoa(i)
	}

	for i := 0; i < 10000; i++ {
		go func() {
			for i := 0; i < 20; i++ {
				a := m[i]
				fmt.Print(a)
			}
		}()
	}

	<-sigs
}
