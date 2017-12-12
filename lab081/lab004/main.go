package main

import (
	"log"
	"time"
)

func main() {
	myTime := "2017-02-08"
	log.Println("myTime=", myTime)

	//t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	t, err := time.Parse("2006-01-02", myTime)
	if err != nil {
		log.Fatalf("time.Parse error:%v", err)
	}
	log.Printf("year=%d,month=%d,day=%d\n", t.Year(), t.Month(), t.Day())
}
