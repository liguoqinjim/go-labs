package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	c := cron.New()
	entryId, err := c.AddFunc("30 * * * *", func() { log.Println("Every hour on the half hour") })
	if err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	}
	c.Start()

	log.Println("下次运行时间:", c.Entry(entryId).Next)
	c.Stop() // Stop the scheduler (does not stop any jobs already running).

	//秒级
	c = cron.New(cron.WithSeconds())
	entryId, err = c.AddFunc("10 * * * * *", func() { log.Println("Event minute on the 10th second") })
	if err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	}
	c.Start()
	log.Println("下次运行时间:", c.Entry(entryId).Next)
	c.Stop()

	//重新启动
	c.Start()
	log.Println("entries:", c.Entries())

	//秒级添加5 fields的
	entryId, err = c.AddFunc("30 * * * *", func() { log.Println("Every hour on the half hour") })
	if err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	}

	c.Stop()
}
