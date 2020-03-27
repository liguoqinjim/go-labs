package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	conn []string
	c    *zk.Conn
)

func init() {
	pflag.StringArrayVarP(&conn, "conn", "c", []string{"127.0.0.1:2181"}, "zookeeper connection")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	demo()

	<-sigs
	c.Close()
	log.Println("program end")
}

func demo() {
	var err error
	var ch <-chan zk.Event
	c, ch, err = zk.Connect(conn, time.Second, zk.WithLogInfo(false))
	if err != nil {
		log.Fatalf("zk.Connet error:%v", err)
	}

	go func() {
		for {
			select {
			case e := <-ch:
				log.Printf("get event:%+v", e)
			}
		}
	}()

	//创建临时node
	//zk.FlagEphemeral表示创建的node是临时节点
	pathEphemeral := "/zoo_temp"
	if r, err := c.Create(pathEphemeral, []byte("123"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err != nil {
		log.Fatalf("c.Create error:%v", err)
	} else {
		log.Printf("create ephemeral node result:%s", r)
	}

	//查询
	if data, _, err := c.Get(pathEphemeral); err != nil {
		log.Fatalf("Get returned error: %+v", err)
	} else {
		log.Printf("get [%s]=%s", pathEphemeral, data)
	}

	//创建永久空node
	path := "/zoo"
	if r, err := c.Create(path, nil, 0, zk.WorldACL(zk.PermAll)); err != nil {
		log.Fatalf("c.Create error:%v", err)
	} else {
		log.Printf("create note result:%s", r)
	}

	//查询空节点
	if data, _, err := c.Get(path); err != nil {
		log.Fatalf("Get returned error: %+v", err)
	} else {
		log.Printf("get [%s]=%s", path, data)
	}

	//加入组
	names := []string{"cat", "dog", "dolphin"}
	for _, name := range names {
		p := path + "/" + name
		if r, err := c.Create(p, []byte(name), 0, zk.WorldACL(zk.PermAll)); err != nil {
			log.Fatalf("c.Create error:%v", err)
		} else {
			log.Printf("create noed result:%s", r)
		}

		if data, _, err := c.Get(p); err != nil {
			log.Fatalf("Get returned error: %+v", err)
		} else {
			log.Printf("get [%s]=%s", p, data)
		}
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
}
