package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("America/Vancouver")
	fmt.Printf("Right now in Vancouver is %s\n", today)

	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	nextOlympics, _ := carbon.CreateFromDate(2016, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	fmt.Printf("Next olympics are in %d\n", nextOlympics.Year())

	if carbon.Now().IsWeekend() {
		fmt.Printf("Party time!")
	}

	//时间戳
	c, err := carbon.CreateFromTimestamp(1603783275, "Asia/Shanghai")
	if err != nil {
		log.Fatalf("CreateFromTimestamp error:%v", err)
	}
	log.Println(c.DateTimeString())

	//得到下一个周三
	log.Println(c.Next(time.Wednesday).DateTimeString())
}
