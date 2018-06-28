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

const TEST_KEY = "key003"
const TEST_VALUE = "value003"

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

	if err != nil {
		log.Printf("clientv3.New error:%v", err)
	} else {
		log.Println("clientv3.New success")
		defer cli.Close()
	}

	//grant
	resp, err := cli.Grant(context.TODO(), 15)
	if err != nil {
		log.Printf("cli.Grant error:%v", err)
	} else {
		log.Printf("cli.Grant success")
		log.Printf("grant.resp:ID=[%d],TTL[%d]", resp.ID, resp.TTL)
	}

	//put
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	rsp, err := cli.Put(ctx, TEST_KEY, TEST_VALUE, clientv3.WithLease(resp.ID))
	if err != nil {
		log.Printf("put error:%v", err)
	} else {
		log.Println("put success")
		log.Printf("rsp.Header=%#v", rsp.Header)
		log.Printf("rsp.PrevKv=%#v", rsp.PrevKv)
	}

	//keepalive
	kpAliveCh, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Printf("cli.KeepAlive error:%v", err)
	} else {
		log.Printf("cli.KeepAlive success")
	}

	//处理keepalive得到的kpAliceCh
	go func() {
		for {
			select {
			case kpResp := <-kpAliveCh:
				log.Printf("kpResp.ID=[%d],TTL=[%d]", kpResp.ID, kpResp.TTL)
			}
		}
	}()

	<-sigs
	log.Println("program end")
}
