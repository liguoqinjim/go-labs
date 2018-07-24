package main

import (
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	//第一个参数是默认的缓存
	//第二个参数是多久去清理一次过期的数据
	//c := cache.New(5*time.Minute, 10*time.Minute)
	c := cache.New(10*time.Second, 15*time.Second)

	go func() {
		//默认缓存时间
		c.Set("foo", "bar", cache.DefaultExpiration)

		v, found := c.Get("foo")
		if found {
			log.Println("key foo found:", v)
		} else {
			log.Println("key foo not found")
		}

		time.Sleep(time.Second * 9)
		v, found = c.Get("foo")
		if found {
			log.Println("key foo found:", v)
		} else {
			log.Println("key foo not found")
		}

		time.Sleep(time.Second * 2)
		v, found = c.Get("foo")
		if found {
			log.Println("key foo found:", v)
		} else {
			log.Println("key foo not found")
		}
	}()

	go func() {
		//key设置为不过时
		c.Set("foo2", "bar2", cache.NoExpiration)

		if v, found := c.Get("foo2"); found {
			s := v.(string)
			log.Println("key foo2 value:", s)
		} else {
			log.Println("key foo2 not found")
		}

		//删除key
		c.Delete("foo2")

		if v, found := c.Get("foo2"); found {
			s := v.(string)
			log.Println("key foo2 value:", s)
		} else {
			log.Println("key foo2 not found")
		}
	}()

	<-sigs
	log.Println("program end")
}
