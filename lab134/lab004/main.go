package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//读取数据，用|隔开
	conf, err := ioutil.ReadFile("zk.conf")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	conns := strings.Split(string(conf), "|")
	c, ech, err := zk.Connect(conns, time.Second)
	if err != nil {
		log.Fatalf("zk.Connect error:%v", err)
	}
	go func() {
		for {
			select {
			case e := <-ech:
				log.Printf("get Event :%+v", e)
			}
		}
	}()

	//创建znode
	path := "/" + "lab004"
	r, err := c.Create(path, []byte("123"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatalf("Create error:%v", err)
	} else {
		log.Println("r=", r)
	}

	<-sigs
	log.Println("program end")
}
