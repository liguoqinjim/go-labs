package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := 1501315889

	myTime := time.Unix(int64(timestamp), 0)
	fmt.Println(myTime.Format("2006-01-02 15:04:05"))
}
