package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
	"time"
)

func simTimingFunction1() { //模拟20毫秒完成的函数
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 2)
	}
	time.Sleep(time.Hour)
}

func main() {
	//开启trace
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//主要逻辑
	for i := 0; i < 100; i++ {
		go simTimingFunction1()
	}

	//使用signal来停止程序
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("exiting")
}
