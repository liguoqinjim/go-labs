package main

import (
	"github.com/shirou/gopsutil/cpu"
	"log"
)

func main() {
	infos, err := cpu.Info()
	if err != nil {
		log.Fatalf("cpu.Info error:%v", err)
	}
	for n, v := range infos {
		log.Println(n, v)
	}

	perfInfos, err := cpu.PerfInfo()
	if err != nil {
		log.Fatalf("perfInfos error:%v", err)
	}
	for n, v := range perfInfos {
		log.Printf("perfInfos[%d]:%+v", n, v)
	}

	procInfos, err := cpu.ProcInfo()
	if err != nil {
		log.Fatalf("procInfos error:%v", err)
	}
	for n, v := range procInfos {
		log.Printf("procInfos[%d]:%+v", n, v)
	}
}
