package main

import (
	"log"
	"time"
)

func main() {
	//demo01()

	demo02()
}

func demo01() {
	myTime := "2017-02-08"
	log.Println("str=", myTime)

	//t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	t, err := time.Parse("2006-01-02", myTime)
	if err != nil {
		log.Fatalf("time.Parse error:%v", err)
	}
	log.Println("t=", t)
	log.Printf("year=%d,month=%d,day=%d\n", t.Year(), t.Month(), t.Day())

	//时区
	t2, err := time.ParseInLocation("2006-01-02", myTime, time.Local)
	log.Println("t2=", t2)
}

//string里面只有时间
func demo02() {
	myTime := "13:14:30"

	t, err := time.Parse("15:04:05", myTime)
	if err != nil {
		log.Fatalf("time.Parse error:%v", err)
	}
	log.Println("t=", t) //t=0000-01-01 13:14:30 +0000 UTC
}
