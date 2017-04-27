package main

import (
	"github.com/funny/snet/go"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	addr := "192.168.116.130:8881"

	config := snet.Config{
		EnableCrypt:        false,
		HandshakeTimeout:   time.Second * 5,
		RewriterBufferSize: 1024,
		ReconnWaitTimeout:  time.Minute * 5,
	}

	conn, err := snet.Dial(config, func() (net.Conn, error) {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Println("net.Dial error", err)
			return nil, err
		}
		return conn, nil
	})
	if err != nil {
		log.Println("conn error", err)
	}

	//发送数据
	go handleConn(conn)
	go listenConn(conn)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("client结束")
}

func handleConn(conn net.Conn) {
	log.Println("发送数据")
	sPos := 0
	moveSpeed := 2

	//每秒发送一次移动数据
	for i := 0; i < 10; i++ {
		pos := sPos + moveSpeed*i
		content := "myPosition" + strconv.Itoa(pos)

		if _, err := conn.Write([]byte(content)); err != nil {
			log.Println("write error", content, err)
		} else {
			log.Println("write success", content)
		}
		time.Sleep(time.Second * 5)
	}
}

func listenConn(conn net.Conn) {
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("read error", err)
			continue
		}

		data := buffer[:n]
		log.Println("收到数据", string(data))
	}
}
