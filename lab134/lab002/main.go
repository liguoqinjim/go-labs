package main

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	conf, err := ioutil.ReadFile("zk.conf")
	if err != nil {
		log.Fatalf("readAll error:%v", err)
	}

	c, ch, err := zk.Connect([]string{string(conf)}, time.Second) //*10)
	if err != nil {
		log.Fatalf("zk.Connet error:%v", err)
	}

	go func() {
		for {
			select {
			case event := <-ch:
				log.Printf("get a event:%+v", event)
			}
		}
	}()

	//创建分组
	path := "/" + "zoo"
	r, err := c.Create(path, []byte("123"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatalf("c.Create error:%v", err)
	}
	log.Println("r=", r)
	data, _, err := c.Get(path)
	if err != nil {
		log.Fatalf("Get returned error: %+v", err)
	}
	log.Printf("data=%s", data)

	//加入组
	names := []string{"cat", "dog", "dolphin"}
	for _, name := range names {
		p := path + "/" + name
		r, err = c.Create(p, []byte(name), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			log.Fatalf("c.Create error:%v", err)
		}
		log.Println("r=", r)
		data, _, err = c.Get(p)
		if err != nil {
			log.Fatalf("Get returned error: %+v", err)
		}
		log.Printf("data=%s", data)
	}

	//成员列表
	children, _, err := c.Children(path)
	if err != nil {
		log.Fatalf("Children error:%v", err)
	}
	log.Println("children=", children)

	//删除分组
	//删除的时候需要版本号，需要版本号也一样才可以删除，但是-1的时候就不用满足这个条件
	for _, name := range names {
		p := path + "/" + name
		err := c.Delete(p, -1)
		if err != nil {
			log.Fatalf("delete error:%v", err)
		} else {
			log.Println("delete success")
		}
	}
	//删除的时候要没有children
	err = c.Delete(path, -1)
	if err != nil {
		log.Fatalf("delete error:%v", err)
	} else {
		log.Println("delete success")
	}

	<-sigs
	log.Println("program end")
}
