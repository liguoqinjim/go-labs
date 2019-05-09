package main

import "log"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		log.Println("received message", msg)
	default:
		log.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		log.Println("send message", msg)
	default:
		//这里因为没人在receive,所以会走default
		log.Println("no message sent")
	}

	select {
	case msg := <-messages:
		log.Println("received message", msg)
	case sig := <-signals:
		log.Println("received signal", sig)
	default:
		log.Println("no activity")
	}
}
