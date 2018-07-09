package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("../nsq.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(string(data), config)
	if err != nil {
		log.Fatal("NewProducer error", err)
	}
	defer w.Stop()

	//发送10条消息
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("hello %d", i)
		err := w.Publish("test_topic", []byte(message))
		if err != nil {
			log.Println("Publish error", err)
		} else {
			log.Println("Publish success ", message)
		}

		time.Sleep(time.Second * 2)
	}
}
