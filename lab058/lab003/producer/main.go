package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("192.168.116.130:4150", config)
	if err != nil {
		log.Fatal("NewProducer error ", err)
	}
	defer w.Stop()

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("hello %d", i)
		err := w.Publish("test003", []byte(message))
		if err != nil {
			log.Fatal("Publish error ", err)
		} else {
			log.Println("Publish ", message)
		}

		time.Sleep(time.Second * 2)
	}
}
