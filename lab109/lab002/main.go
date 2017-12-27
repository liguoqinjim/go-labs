package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func fakeReadDB() {
	ms := make([]int, 100)
	for n := range ms {
		ms[n] = n
	}

	total := 0
	for _, v := range ms {
		total += v
	}
}

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("os.Create error:%v", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 10000; i++ {
		fakeReadDB()
	}
}
