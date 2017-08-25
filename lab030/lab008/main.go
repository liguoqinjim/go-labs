package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
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
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

	time.Sleep(time.Second * 3)
}

func sender(conn net.Conn) {
	words := "hello world!"
	conn.Write([]byte(words))
	fmt.Println("send over")
}
