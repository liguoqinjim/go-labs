package main

import (
	"fmt"
	"github.com/funny/snet/go"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	addr := "0.0.0.0:8881"

	config := snet.Config{
		EnableCrypt:        false,
		HandshakeTimeout:   time.Second * 5,
		RewriterBufferSize: 1024,
		ReconnWaitTimeout:  time.Minute * 5,
	}

	conn, err := snet.Dial(config, func() (net.Conn, error) {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}
		return conn, nil
	})
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 1; i++ {
		content := "content" + strconv.Itoa(i)
		data := []byte(content)

		if _, err := conn.Write(data); err != nil {
			log.Println("write failed", err)
			return
		}
		fmt.Println("发送数据", content)

		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("read failed", err)
			return
		}
		result := buffer[:n]
		fmt.Println("收到返回", string(result))

		//sleep
		time.Sleep(time.Second * 2)
	}
}
