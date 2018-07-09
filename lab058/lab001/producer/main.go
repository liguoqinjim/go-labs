package main

import (
	"log"

	"github.com/nsqio/go-nsq"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("../nsq.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(string(data), config)

	err = w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
