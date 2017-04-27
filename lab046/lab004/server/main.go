package main

import (
	"github.com/funny/snet/go"
	"log"
	"net"
	"os"
	"os/signal"
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

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
