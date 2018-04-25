package main

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("zk.conf")
	if err != nil {
		log.Fatalf("readAll error:%v", err)
	}

	c, ch, err := zk.Connect([]string{string(data)}, time.Second) //*10)
	if err != nil {
		log.Fatalf("zk.Connet error:%v", err)
	}
	_ = c
	_ = ch
}
