package main

import (
	"github.com/google/gops/agent"
	"log"
	"time"
)

func main() {
	err := agent.Listen(agent.Options{
		ShutdownCleanup: true,
	})
	if err != nil {
		log.Fatalf("agent.Listen error:%v", err)
	}

	time.Sleep(time.Hour)
}
