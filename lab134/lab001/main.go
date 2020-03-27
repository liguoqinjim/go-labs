package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"log"
)

var (
	conn []string
	test bool
)

func init() {
	pflag.StringArrayVarP(&conn, "conn", "c", []string{"127.0.0.1:2181"}, "zookeeper connection")
	pflag.BoolVarP(&test, "test", "t", false, "test mode")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	//连接zk
	var c *zk.Conn
	var err error

	if !test {
		c, _, err = zk.Connect(conn, time.Second)
	} else {
		//不输出zk本身的日志
		c, _, err = zk.Connect(conn, time.Second, zk.WithLogInfo(false))
	}
	if err != nil {
		log.Fatalf("zk.Connect error:%v", err)
	}
	defer c.Close()

	//watch children
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		log.Fatalf("ChildrenW error:%v", err)
	}

	log.Printf("children:%+v", children)
	log.Printf("stat:%+v", stat)
	e := <-ch
	log.Printf("ch evnet:%+v", e)
}
