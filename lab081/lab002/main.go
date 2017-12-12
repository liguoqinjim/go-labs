package main

import (
	"log"
	"time"
)

func main() {
	timestamp := 1501315889

	myTime := time.Unix(int64(timestamp), 0)
	log.Println("format1=", myTime.Format("2006-01-02 15:04:05"))

	log.Println("format2=", myTime.Format("20060102"))
}
