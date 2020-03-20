package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	c := cron.New()
	if err := c.AddFunc("0 0/5 * * * ? ", func() {
		log.Println("Every 5 minutes")
	}); err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	}

	if err := c.AddFunc("@hourly", func() {
		log.Println("Every hour")
	}); err != nil {
		log.Fatalf("c.AddFunc error:%v", err)
	}
	c.Start()

	time.Sleep(time.Hour)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
