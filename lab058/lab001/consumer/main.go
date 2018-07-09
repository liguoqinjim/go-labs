package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
	"io/ioutil"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	data, err := ioutil.ReadFile("nqs.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)
		wg.Done()
		return nil
	}))
	err = q.ConnectToNSQD(string(data))
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
