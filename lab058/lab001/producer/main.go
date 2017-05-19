package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("192.168.116.130:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
