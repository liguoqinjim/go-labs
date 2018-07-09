package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//读取ip.conf
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("readfile error:%v", err)
	}

	//连接
	config := nsq.NewConfig()
	w, err := nsq.NewProducer(string(data), config)
	if err != nil {
		log.Fatalf("nsq.NewProducer:%v", err)
	}
	defer w.Stop()

	//发送消息
	i := 0
	for {
		message := fmt.Sprintf("hello %d", i)

		if i%2 == 0 {
			if err := w.Publish("test1", []byte(message)); err != nil {
				log.Println("发送失败1")
			} else {
				log.Println("发送消息1", message)
			}
		} else {
			if err := w.Publish("test2", []byte(message)); err != nil {
				log.Println("发送失败2")
			} else {
				log.Println("发送消息2", message)
			}
		}

		i++

		time.Sleep(time.Second * 3)
	}

	//结束
	<-sigs
	log.Println("end")
}
