package main

import (
	"log"
	"net"
	"time"
)

func main() {
	dialer := &net.Dialer{
		Timeout: time.Second * 2,
	}

	conn, err := dialer.Dial("tcp", "ubuntu:4150")
	if err != nil {
		log.Fatal("dial error=", err)
	}
	defer conn.Close()
}
