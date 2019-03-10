package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("0 0/5 * * * ? ", func() {
		log.Println("Every 5 minutes")
	})
	c.AddFunc("@hourly", func() {
		log.Println("Every hour")
	})
	c.Start()

	time.Sleep(time.Hour)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
