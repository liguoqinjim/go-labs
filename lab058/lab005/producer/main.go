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
	defer w.Stop()

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("hello %d", i)
		err := w.Publish("test005", []byte(message))
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
