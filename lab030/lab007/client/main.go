package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	addr := "localhost:8881"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	//发送数据
	for i := 0; i < 10; i++ {
		content := "helloworld" + strconv.Itoa(i)
		n, err := conn.Write([]byte(content))
		if err != nil {
			log.Println("发送数据错误", err)
		} else {
			log.Printf("发送数据成功第%d次,数据长度%d\n", i, n)
		}

		time.Sleep(time.Second * 3)
	}
}
