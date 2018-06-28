package main

import (
	"github.com/coreos/etcd/clientv3"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	confData, err := ioutil.ReadFile("etcd.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	log.Println("开始时间:", time.Now())
	ips := strings.Split(string(confData), "|")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ips,
		DialTimeout: time.Second * 5, //连接超时时间
	})

	log.Println("结束时间:", time.Now())
	if err != nil {
		log.Printf("clientv3.New error:%v", err)
	} else {
		//close要放在error为空的时候
		defer cli.Close()

		log.Println("clientv3.New success")
	}
}
