package main

import (
	"log"
	"time"
)

func main() {
	//当前时间
	timeNow := time.Now()
	//second
	second := timeNow.Unix()
	log.Println("second=", second)
	//milliSecond
	milliSecond := second * (int64(time.Second) / int64(time.Millisecond))
	log.Println("mill=", milliSecond)
	//microSecond
	microSecond := milliSecond * (int64(time.Millisecond) / int64(time.Microsecond))
	log.Println("micro=", microSecond)
	//nanoSecond
	nanoSecond := microSecond * (int64(time.Microsecond) / int64(time.Nanosecond))
	log.Println("nano=", nanoSecond)
}
