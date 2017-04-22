package main

import (
	"fmt"
	"github.com/kavu/go_reuseport"
)

const addr = "localhost:8881"

func main() {
	listener, err := reuseport.Listen("tcp", "localhost:8881")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("server2开始监听", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("server2收到数据")
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := buffer[:n]
		fmt.Print(string(result))
		conn.Write([]byte("this is server2\n"))
		conn.Close()
	}
}
