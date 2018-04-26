package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//读取数据，用|隔开
	conf, err := ioutil.ReadFile("../zk.conf")
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

	//创建节点
	path := "/" + "lab005"
	r, err := c.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatalf("Create error:%v", err)
	} else {
		log.Println("r=", r)
	}

	//修改节点值
	for i := 1; i <= 7; i++ {
		v := i * 3
		_, err := c.Set(path, []byte(strconv.Itoa(v)), -1)
		if err != nil {
			log.Printf("Set error:%v", err)
		} else {
			log.Printf("Set success:%d", v)
		}

		time.Sleep(time.Second * 3)
	}

	<-sigs
	log.Println("master program end")
}
