package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	c := cron.New()

	if err := c.AddFunc("0 6 * * * ?", func() {
		log.Println(1)
	}); err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	} else {
		log.Println("start1")
		c.Start()
		log.Println("start2")
	}
	log.Println("next:", c.Entries()[0].Next)

	//添加秒级的cron
	if err := c.AddFunc("0/10 * * * * ? *", func() {
		log.Println(1)
	}); err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	}else{
		log.Println("秒级cron success")
	}

	time.Sleep(time.Hour)
}
