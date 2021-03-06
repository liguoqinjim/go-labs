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
	log.Printf("second=\t%d", second)
	//milliSecond
	milliSecond := second * (int64(time.Second) / int64(time.Millisecond))
	log.Printf("mill=\t%d", milliSecond)
	//microSecond
	microSecond := milliSecond * (int64(time.Millisecond) / int64(time.Microsecond))
	log.Printf("micro=\t%d", microSecond)
	//nanoSecond
	nanoSecond := microSecond * (int64(time.Microsecond) / int64(time.Nanosecond))
	log.Printf("nano=\t%d", nanoSecond)
}
