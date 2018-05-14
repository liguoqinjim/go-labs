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

const KEY_NAME_PREFIX = "lab005"

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

	//先put一组值
	for i := 0; i < 3; i++ {
		_, err := cli.Put(context.TODO(), KEY_NAME_PREFIX+"/"+strconv.Itoa(i), "sample"+strconv.Itoa(i))
		if err != nil {
			log.Printf("cli.Put error:%v", err)
		} else {
			log.Printf("cli.Put success:%d", i)
		}
	}

	ch := cli.Watch(context.TODO(), KEY_NAME_PREFIX+"/", clientv3.WithPrefix())
	go func() {
		for {
			log.Printf("receive...")
			select {
			case c := <-ch:
				for _, e := range c.Events {
					log.Printf("get event:%+v", e)
				}
			}
		}
	}()

	<-sigs
	log.Println("program end")
}
