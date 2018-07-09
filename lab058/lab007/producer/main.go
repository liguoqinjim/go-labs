package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(string(data), config)
	if err != nil {
		log.Fatal("NewProducer error:", err)
	}

	for i := 0; i < 100; i++ {
		message := fmt.Sprintf("hello %d", i)
		err := w.Publish("lab007", []byte(message))
		if err != nil {
			log.Fatal("Publish error:", err)
		} else {
			log.Println("发送消息:", message)
		}

		time.Sleep(time.Second * 2)
	}
}
