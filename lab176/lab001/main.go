package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	//每秒100个，输出的就是10ms一个 10ms*100=1s，最小就是1个
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
