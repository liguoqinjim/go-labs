package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
