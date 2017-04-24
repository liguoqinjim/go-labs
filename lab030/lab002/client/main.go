package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverAddr := "localhost:8881"

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("tcpAddr=", tcpAddr)

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}

	words := "hello world!"
	conn.Write([]byte(words))
	fmt.Println("send over")
	conn.Close()
}
