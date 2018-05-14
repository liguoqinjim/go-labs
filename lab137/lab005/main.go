package main

import (
	"github.com/coreos/etcd/clientv3"
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

	confData, err := ioutil.ReadFile("etcd.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	ips := strings.Split(string(confData), "|")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ips,
		DialTimeout: time.Second * 2,
	})
	defer cli.Close()

	if err != nil {
		log.Printf("clientv3.New error:%v", err)
	} else {
		log.Println("clientv3.New success")
	}

	<-sigs
	log.Println("program end")
}
