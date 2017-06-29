package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("test003", "ch", config)
	if err != nil {
		log.Fatal("NewConsumer error ", err)
	}
	q.AddHandler(nsq.HandlerFunc(MsgHandler))

	err = q.ConnectToNSQLookupd("192.168.116.130:4161")
	if err != nil {
		log.Fatal("ConnectToNSQLookupd errro ", err)
	}

	<-sigs
	q.Stop()
}

func MsgHandler(message *nsq.Message) error {
	log.Println("收到消息", string(message.ID[:]), string(message.Body))
	return nil
}
