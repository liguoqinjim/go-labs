package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", ":3333")
	CheckError(err)
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		log.Println(conn.RemoteAddr().String(), " tcp connect success")
		handleConnnection(conn)
	}
}

func handleConnnection(conn net.Conn) {
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			log.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		log.Println(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s", err.Error())
		os.Exit(1)
	}
}
