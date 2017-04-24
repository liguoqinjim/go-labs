package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const (
	HEART_CONTENT = "h" //心跳包信息
)

func main() {
	addr := "localhost:8881"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}

	go sendHeart(conn, 2500)
	send(conn)
}

func send(conn net.Conn) {
	for i := 0; i < 5; i++ {
		content := "content" + strconv.Itoa(i)
		conn.Write([]byte(content))
		fmt.Println("发送数据", content)

		interval := rand.Intn(6) + 1 //随机暂停
		time.Sleep(time.Second * time.Duration(interval))
	}
}

func sendHeart(conn net.Conn, heartInterval int) { //发送心跳信息,heartInterval是心跳信号的间隔(毫秒)
	ticker := time.NewTicker(time.Millisecond * time.Duration(heartInterval))
	for range ticker.C {
		conn.Write([]byte(HEART_CONTENT))
		fmt.Println("发送心跳信号", time.Now())
	}
}
