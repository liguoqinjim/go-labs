package main

import (
	"github.com/agatan/timejump"
	"log"
	"time"
)

func main() {
	timejump.Activate()
	defer timejump.Deactivate()

	timejump.Stop()
	timejump.Jump(time.Date(1992, time.November, 15, 0, 0, 0, 0, time.Local))

	now := timejump.Now()
	log.Println("Year =", now.Year())
	log.Println("Now =", now)
}
