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

	c, _, err := zk.Connect([]string{string(data)}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	log.Printf("%+v %+v\n", children, stat)
	e := <-ch
	log.Printf("%+v\n", e)
}
