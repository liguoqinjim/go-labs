package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	CONN_TIMEOUT = 3 //socket超时
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8881")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}

		handleTimeout(conn)
		data := buffer[:n]

		content := string(data)
		if content == "h" {
			fmt.Println("收到心跳信号")
		} else {
			fmt.Println("收到消息", string(data))
		}
	}
}

func handleTimeout(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(time.Second * time.Duration(CONN_TIMEOUT)))
}
