package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

type Student struct {
	Sid   int
	Sname string
	Sage  int
}

func main() {
	time1 := time.Now().Unix()
	fmt.Println("起点:", time.Now())
	//监控cpu
	f, err := os.OpenFile("cpu.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("err=", err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	fmt.Println("监控起点:", time.Now())
	defer pprof.StopCPUProfile()

	fmt.Println("hello lab008")
	for i := 0; i < 2000; i++ {
		time.Sleep(time.Millisecond * 10)
	}

	s := &Student{Sid: 1, Sname: "xiaoming", Sage: 20}
	fmt.Println(s)

	time2 := time.Now().Unix()
	fmt.Println("终点:", time.Now())
	total_time := time2 - time1
	fmt.Println("总共时间:", total_time)
}
