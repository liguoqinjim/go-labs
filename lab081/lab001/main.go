package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()

	log.Println("now=", now)
	log.Println("now.Unix=", now.Unix())
	log.Println("now.UnixNano=", now.UnixNano())

	log.Println("format1=", now.Format("2006-01-02"))
	log.Println("format2=", now.Format("2006-01-02 15:04:05"))
}
