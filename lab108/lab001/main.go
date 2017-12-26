package main

import (
	"github.com/shirou/gopsutil/mem"
	"log"
)

func main() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("mem.VirtualMemory error:%v", err)
	}

	log.Printf("Total: %v, Free:%v, UsedPercent:%f%%", v.Total, v.Free, v.UsedPercent)

	log.Println(v)
}
