package main

import (
	"github.com/shirou/gopsutil/host"
	"log"
)

func main() {
	info, err := host.Info()
	if err != nil {
		log.Fatalf("host.Info error:%v", err)
	}
	log.Println("info=", info)

	osInfo, err := host.GetOSInfo()
	if err != nil {
		log.Fatalf("host.GetOSInfo error:%v", err)
	}
	log.Printf("osInfo=%+v", osInfo)
}
