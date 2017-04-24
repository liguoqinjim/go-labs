package main

import (
	"fmt"
	"lab030/lab004/protocol"
	"log"
	"net"
)

func main() {
	listerer, err := net.Listen("tcp", "localhost:8881")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listerer.Accept()
		if err != nil {
			continue
		}

		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		pack := buffer[:n]
		data, err := protocol.Depack(pack)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("收到消息:", string(data))
	}
}
