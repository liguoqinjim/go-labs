package main

import (
	"fmt"
	"github.com/funny/snet/go"
	"log"
	"net"
	"time"
)

func main() {
	config := snet.Config{
		EnableCrypt:        false,
		HandshakeTimeout:   time.Second * 5,
		RewriterBufferSize: 1024,
		ReconnWaitTimeout:  time.Minute * 5,
	}

	listener, err := snet.Listen(config, func() (net.Listener, error) {
		l, err := net.Listen("tcp", "0.0.0.0:8881")
		checkErr(err)
		return l, err
	})
	checkErr(err)

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
			log.Println(err)
			return
		}
		content := buffer[:n]
		fmt.Println("服务器收到消息:content=", string(content))

		//返回客户端
		data := string(content) + "<-server"
		conn.Write([]byte(data))

		//模拟断开连接
		if i == 3{

		}
		conn.Close()
		i++
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
