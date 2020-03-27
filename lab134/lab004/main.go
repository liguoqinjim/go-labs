package main

import (
	"flag"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
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

	//连接集群
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

	//创建znode
	path := "/lab004"
	if r, err := c.Create(path, []byte("123"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err != nil {
		log.Fatalf("Create error:%v", err)
	} else {
		log.Printf("create node result:%s", r)
	}

	//znode是否存在
	if exist, stat, err := c.Exists(path); err != nil {
		log.Fatalf("Exists error:%v", err)
	} else {
		log.Printf("exist result:%t,%+v", exist, stat)
	}

	//观察模式
	//是否存在 watch
	exist, stat, watchChan, err := c.ExistsW(path)
	if err != nil {
		log.Fatalf("ExistsW error:%v", err)
	} else {
		log.Printf("existsw result:%t,%+v", exist, stat)
	}

	//监听ch
	go func() {
		//接收第一次
		e := <-watchChan
		log.Printf("exists watch channel event:%+v", e)

		//创建新的watcher,因为watcher只能使用一次
		exist, stat, watchChan2, err := c.ExistsW(path)
		if err != nil {
			log.Fatalf("ExistsW error:%v", err)
		} else {
			log.Printf("existsw result:%t,%+v", exist, stat)
		}

		event2 := <-watchChan2
		log.Printf("exists watch channel event:%+v", event2)
	}()

	<-sigs
	c.Close()
	log.Println("program end")
}
