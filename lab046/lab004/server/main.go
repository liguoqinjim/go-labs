package main

import (
	"github.com/funny/snet/go"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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
		l, err := net.Listen("tcp", "127.0.0.1:8881")
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
		go sender(conn)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("server结束")
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
		log.Println("收到消息", string(content))

		//模拟断开连接
		//if i == 3 {
		//	log.Println("server端模拟连接关闭")
		//	conn.Close()
		//	break
		//}
		i++
	}
}

func sender(conn net.Conn) {
	blood := 100
	for i := 0; i < 10; i++ {
		blood -= rand.Intn(8) + 1
		content := "blood:" + strconv.Itoa(blood)

		if _, err := conn.Write([]byte(content)); err != nil {
			log.Println("write err", err)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
