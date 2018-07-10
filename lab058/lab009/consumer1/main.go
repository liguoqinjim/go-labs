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
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	//创建连接
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("server900001", "ch", config)
	if err != nil {
		log.Fatalf("NewConsumer error:%v", err)
	}
	q.AddHandler(nsq.HandlerFunc(MsgHandler))

	//连接
	err = q.ConnectToNSQLookupd(string(data))
	if err != nil {
		log.Fatalf("ConnnectToNSQLookupd error:%v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Println("end")
}

func MsgHandler(message *nsq.Message) error {
	log.Println("消息处理:", string(message.ID[:]), string(message.Body))
	return nil
}
