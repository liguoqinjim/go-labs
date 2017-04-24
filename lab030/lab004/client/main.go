package main

import (
	"fmt"
	"lab030/lab004/protocol"
	"log"
	"net"
)

func main() {
	addr := "localhost:8881"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}

	sender(conn)
}

func sender(conn net.Conn) {
	defer conn.Close()

	content := "20170424网易云音乐正在播放那就这样吧"
	data := protocol.Enpack([]byte(content))
	fmt.Println("发送协议长度:", len(data))
	conn.Write(data)
}
