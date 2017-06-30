package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	address := readConf()
	log.Println("address=", address)

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(address, config)
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

func readConf() string {
	file, err := os.Open("ip.conf")
	if err != nil {
		log.Fatal("open file error:", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("readAll error:", err)
	}

	return string(data)
}
