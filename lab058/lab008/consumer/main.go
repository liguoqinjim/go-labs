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

	addr := string(data)

	//创建连接
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("test1", "ch", config)
	if err != nil {
		log.Fatalf("NewConsumer error:%v", err)
	}
	q.AddHandler(nsq.HandlerFunc(MsgHandler))

	//连接
	err = q.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Fatalf("ConnnectToNSQLookupd error:%v", err)
	}

	//创建连接2
	q2, err := nsq.NewConsumer("test2", "ch", config)
	if err != nil {
		log.Fatalf("NewConsumer2 error:%v", err)
	}
	q2.AddHandler(nsq.HandlerFunc(Msg2Handler))

	//连接
	err = q2.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Fatalf("ConnectToNSQLookupd2 error:%v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Println("end")
}

func MsgHandler(message *nsq.Message) error {
	log.Println("消息处理1:", string(message.ID[:]), string(message.Body))
	return nil
}

func Msg2Handler(message *nsq.Message) error {
	log.Println("消息处理2:", string(message.ID[:]), string(message.Body))
	return nil
}
