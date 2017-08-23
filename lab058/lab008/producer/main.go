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
	//读取ip.conf
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("readfile error:%v", err)
	}
	addr := string(data)

	//连接
	config := nsq.NewConfig()
	w, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Fatalf("nsq.NewProducer:%v", err)
	}
	defer w.Stop()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//发送消息
	i := 0
	for {
		message := fmt.Sprint("hello %d", i)
		if err := w.Publish("test1", []byte(message)); err != nil {
			log.Println("发送失败")
		} else {
			log.Println("发送消息", message)
		}

		i++

		time.Sleep(time.Second * 3)
	}

	//结束
	<-sigs
	log.Println("end")
}
