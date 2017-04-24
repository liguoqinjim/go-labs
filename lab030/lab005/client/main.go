package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"strconv"
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
	fmt.Println("连接success")
	sender(conn)
}
func sender(conn net.Conn) {
	for i := 0; i < 5; i++ {
		words := strconv.Itoa(i) + "This is a test for long conn"
		fmt.Println("发送数据")
		conn.Write([]byte(words))
		time.Sleep(6 * time.Second)
	}
	fmt.Println("send over")
}
