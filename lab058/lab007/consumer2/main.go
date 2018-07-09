package main

import (
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("lab007", "b", config)
	if err != nil {
		log.Fatal("NewConsumer error:", err)
	}
	q.AddHandler(nsq.HandlerFunc(MsgHandler))

	err = q.ConnectToNSQLookupd(string(data))
	if err != nil {
		log.Fatal("ConnectTONSQLookup error:", err)
	}

	<-sigs
	log.Println("程序结束")
}

func MsgHandler(message *nsq.Message) error {
	log.Println("收到消息", string(message.ID[:]), string(message.Body))
	return nil
}
