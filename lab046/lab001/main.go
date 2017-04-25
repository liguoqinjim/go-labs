package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/funny/snet/go"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	//参数
	encrypt := false
	reconn := false

	//启动服务器
	config := snet.Config{
		EnableCrypt:        false,
		HandshakeTimeout:   time.Second * 5,
		RewriterBufferSize: 1024,
		ReconnWaitTimeout:  time.Minute * 5,
	}

	listener, err := snet.Listen(config, func() (net.Listener, error) {
		l, err := net.Listen("tcp", "0.0.0.0:0")
		if err != nil {
			return nil, err
		}
		return l, nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept failed: %s", err.Error())
			return
		}

		fmt.Println("io.Copy")
		io.Copy(conn, conn)
		conn.Close()
		log.Println("copy exit")
		wg.Done()
	}()

	//客户端
	conn, err := snet.Dial(config, func() (net.Conn, error) {
		conn, err := net.Dial("tcp", listener.Addr().String())
		if err != nil {
			return nil, err
		}
		return conn, nil
	})
	defer conn.Close()

	for i := 0; i < 2; i++ {
		b := RandBytes(100)
		c := b
		if encrypt {
			c = make([]byte, len(b))
			copy(c, b)
		}

		if _, err := conn.Write(b); err != nil {
			log.Fatalf("write failed: %s", err.Error())
			return
		}

		if reconn && i%100 == 0 {
			conn.(*snet.Conn).TryReconn()
		}

		a := make([]byte, len(b))
		if _, err := io.ReadFull(conn, a); err != nil {
			log.Fatalf("read failed: %s", err.Error())
			return
		}

		log.Println(b)
		if !bytes.Equal(a, c) {
			println("i =", i)
			println("a =", hex.EncodeToString(a))
			println("c =", hex.EncodeToString(c))
			log.Fatalf("a != c")
			return
		}
	}

	conn.Close()
	listener.Close()
	wg.Wait()
}

func RandBytes(n int) []byte {
	n = rand.Intn(n) + 1
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}
