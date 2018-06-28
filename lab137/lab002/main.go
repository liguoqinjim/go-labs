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

const KEY_NAME = "lab002"

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	confData, err := ioutil.ReadFile("etcd.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	//连接etcd
	ips := strings.Split(string(confData), "|")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ips,
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		log.Printf("clientv3.New error:%v", err)
	} else {
		log.Println("clientv3.New success")
		defer cli.Close()
	}

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	resp, err := cli.Put(ctx, "lab002", "sample_003")
	cancel()
	if err != nil {
		log.Printf("put error:%v", err)
	} else {
		log.Printf("resp.Header=%#v", resp.Header)
		log.Printf("resp.PrevKv=%#v", resp.PrevKv)
		log.Println("put success")
	}

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	resp2, err := cli.Get(ctx, KEY_NAME)
	cancel()
	if err != nil {
		log.Printf("get error:%v", err)
	} else {
		log.Printf("resp2.Header=%+v", resp2.Header)
		log.Printf("resp2.Kvs=%+v", resp2.Kvs)
		log.Println("get success")
	}

	//del
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	resp3, err := cli.Delete(ctx, KEY_NAME)
	cancel()
	if err != nil {
		log.Printf("del error:%v", err)
	} else {
		log.Printf("resp3.Header=%#v", resp3.Header)
		log.Printf("resp3.Kvs=%#v", resp3.PrevKvs)
		log.Println("del success")
	}

	<-sigs
	log.Println("program ending")
}
