package main

import (
	"flag"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
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

	var err error
	var ch <-chan zk.Event
	c, ch, err = zk.Connect(conn, time.Second, zk.WithLogInfo(false))
	if err != nil {
		log.Fatalf("zk.Connect error:%v", err)
	}
	go func() {
		for {
			select {
			case e := <-ch:
				log.Printf("get Event :%+v", e)
			}
		}
	}()

	//创建节点
	path := "/lab005"
	if r, err := c.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err != nil {
		log.Fatalf("c.Create error:%v", err)
	} else {
		log.Printf("create result:%s", r)
	}

	//修改节点值
	for i := 1; i <= 7; i++ {
		v := i * 3
		if stat, err := c.Set(path, []byte(strconv.Itoa(v)), -1); err != nil {
			log.Fatalf("Set error:%v", err)
		} else {
			log.Printf("set result:%+v", stat)
		}

		time.Sleep(time.Second * 3)
	}

	<-sigs
	c.Close()
	log.Println("master program end")
}
