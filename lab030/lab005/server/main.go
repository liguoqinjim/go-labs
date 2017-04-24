package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8881")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		handleConnection(conn, 2)
	}
}

//长连接入口
func handleConnection(conn net.Conn, timeout int) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Data := (buffer[:n])
		messnager := make(chan byte)
		//心跳计时
		go HeartBeating(conn, messnager, timeout)
		//检测每次Client是否有数据传来
		go GravelChannel(Data, messnager)
		Log("receive data length:", n)
		Log(conn.RemoteAddr().String(), "receive2 data string:", string(Data))
	}
}

//心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func HeartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case fk := <-readerChannel:
		Log(conn.RemoteAddr().String(), "receive1 data string:", string(fk))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		//conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
		break
	case <-time.After(time.Millisecond * 100):
		Log("It's really weird to get Nothing!!!")
		conn.Close()
		break
	}
}

func GravelChannel(n []byte, mess chan byte) {
	for _, v := range n {
		// 这里会有channel堵塞的情况
		mess <- v
	}
	close(mess)
}

func Log(v ...interface{}) {
	log.Println(v...)
}
