package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server := "127.0.0.1:3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
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
