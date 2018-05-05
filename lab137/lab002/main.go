package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	confData, err := ioutil.ReadFile("etcd.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error", err)
	}

	ips := strings.Split(string(confData), "|")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ips,
		DialTimeout: time.Second * 2,
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
		log.Printf("%#v", resp.Header)
		log.Printf("%#v", resp.PrevKv)
		log.Println("put success")
	}

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	resp2, err := cli.Get(ctx, "lab002")
	cancel()
	if err != nil {
		log.Printf("get error:%", err)
	} else {
		log.Printf("%+v", resp2.Header)
		log.Printf("%+v", resp2.Kvs)
		log.Println("get success")
	}
}
