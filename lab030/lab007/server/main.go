package main

import (
	"log"
	"net"
)

func main() {
	addr := "localhost:8881"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 2048)

	i := 0
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("接收数据错误", err)
			return
		}

		log.Printf("接收数据第%d次,数据%s\n", i, string(buffer[:n]))

		if i == 5 {
			log.Printf("模拟连接断开,第%d次\n", i)
			conn.Close()
			break
		}
		i++
	}
}
