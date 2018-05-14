package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const KEY_NAME = "lab003"

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

	//注册服务
	resp, _ := cli.Grant(context.TODO(), 15)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	rsp, err := cli.Put(ctx, KEY_NAME, `{"addr":"192.168.1.1:9999"}`, clientv3.WithLease(resp.ID))
	if err != nil {
		log.Printf("put error:%v", err)
	} else {
		log.Printf("rsp.Header=%#v", rsp.Header)
		log.Printf("rsp.PrevKv=%#v", rsp.PrevKv)
		log.Println("put success")
	}

	<-sigs
	log.Println("program end")
}
