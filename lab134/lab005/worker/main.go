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

	path := "/lab005"

DONE:
	for {
		data, stat, eventChan, err := c.GetW(path)
		if err != nil {
			log.Fatalf("GetW error:%v", err)
		} else {
			log.Printf("GetW data:%s,%+v", data, stat)
		}

		select {
		case e := <-eventChan:
			log.Printf("GetW event:%+v", e)
		case <-sigs:
			break DONE
		}
	}

	log.Println("worker program end")
}
