package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//读取ip
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("reafFile error:%v", err)
	}
	addr := string(data)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr error:%v", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("DialTCP error:%v", err)
	} else {
		log.Println("connect success")
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		i := 0
		for {
			words := fmt.Sprintf("helloworld%d\n", i)

			conn.Write([]byte(words))
			log.Println("send:", words)

			i++
			time.Sleep(time.Second * 5)
		}
	}()

	<-sigs
	log.Println("end")
}
