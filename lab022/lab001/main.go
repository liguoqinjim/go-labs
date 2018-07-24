package main

import (
	"log"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	log.Println("msg=", msg)
}
