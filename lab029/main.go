package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("sig=", sig)
		done <- true
	}()

	fmt.Println("awaiting signal", time.Now())
	<-done
	fmt.Println("exiting", time.Now())
}
