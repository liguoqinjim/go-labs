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
	r, err := c.Create(path, []byte("123"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatalf("Create error:%v", err)
	} else {
		log.Println("r=", r)
	}

	//是否存在
	e, s, err := c.Exists(path)
	if err != nil {
		log.Fatalf("Exists error:%v", err)
	} else {
		log.Println("exists=", e, s)
	}

	//是否存在 watch
	e, s, ew, err := c.ExistsW(path)
	if err != nil {
		log.Fatalf("ExistsW error:%v", err)
	} else {
		log.Println("existsw=", e, s)
	}
	go func() {
		//接收第一次
		e := <-ew
		log.Printf("exists watch:%+v", e)

		//创建新的watcher,因为watcher只能使用一次
		e2, s2, ew2, err2 := c.ExistsW(path)
		if err2 != nil {
			log.Fatalf("ExistsW error2:%+v", err2)
		} else {
			log.Println("exists2=", e2, s2)
		}

		event2 := <-ew2
		log.Printf("exists watch2:%+v", event2)
	}()

	<-sigs
	log.Println("program end")
}
