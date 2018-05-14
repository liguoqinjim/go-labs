package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const KEY_NAME_PREFIX = "lab004"

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

	//put多个值
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		_, err := cli.Put(ctx, KEY_NAME_PREFIX+"/key"+strconv.Itoa(i), "sample"+strconv.Itoa(i))
		cancel()
		if err != nil {
			log.Printf("put error:%v", err)
		} else {
			log.Printf("put success:%d", i)
		}
	}

	//get多个值
	rsp, err := cli.Get(context.TODO(), KEY_NAME_PREFIX+"/", clientv3.WithPrefix())
	if err != nil {
		log.Printf("cli.Get error:%v", err)
	} else {
		log.Printf("rsp.Header=%+v", rsp.Header)
		log.Printf("rsp.Kvs=%+v", rsp.Kvs)
		for _, v := range rsp.Kvs {
			log.Printf("key[%s]value[%s]", v.Key, v.Value)
		}
		log.Println("get success")
	}

	//删除多个值
	resp2, err := cli.Delete(context.TODO(), KEY_NAME_PREFIX+"/", clientv3.WithPrefix())
	if err != nil {
		log.Printf("delete error:%v", err)
	} else {
		log.Printf("resp2=%+v", resp2.Header)
		log.Printf("resp2=%+v", resp2.PrevKvs)
		log.Println("delete success")
	}

	<-sigs
	log.Println("program end")
}
