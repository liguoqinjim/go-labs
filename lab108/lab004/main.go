package main

import (
	"github.com/shirou/gopsutil/load"
	"log"
)

func main() {
	avg, err := load.Avg()
	if err != nil {
		log.Println("load.Avg error:%v", err)
	}
	log.Println("avg=", avg)

	misc, err := load.Misc()
	if err != nil {
		log.Println("load.Misc error:%v", err)
	}
	log.Println("misc=", misc)
}
