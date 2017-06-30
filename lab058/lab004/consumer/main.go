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
	ip := readConf()
	log.Println("nsq address=", ip)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("test004", "ch", config)
	if err != nil {
		log.Fatal("NewConsumer error:", err)
	}
	q.AddHandler(nsq.HandlerFunc(MsgHandler))

	err = q.ConnectToNSQLookupd(ip)
	if err != nil {
		log.Fatal("ConnectToNSQLookupd error:", err)
	}

	<-sigs
	log.Println("程序结束")
}

func MsgHandler(message *nsq.Message) error {
	log.Println("收到消息:", string(message.ID[:]), string(message.Body))
	return nil
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
