package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8881")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
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
			return
		}

		fmt.Println("收到消息", string(buffer[:n]))
	}
}
