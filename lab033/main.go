package main

import (
	"github.com/google/gops/agent"
	"log"
	"time"
)

func main() {
	if err := agent.Listen(nil); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
