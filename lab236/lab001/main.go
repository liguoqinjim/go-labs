package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func task() {
	fmt.Println("I am running task.", time.Now())
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func taskHour() {
	fmt.Println("taskHour")
}

func main() {
	//每三秒
	gocron.Every(3).Second().Do(task)
	//带参数
	gocron.Every(3).Second().Do(taskWithParams, 1, "a")

	//指定开始时间为下一次周期，也可以是别的时间
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)

	// NextRun gets the next running time
	_, t := gocron.NextRun()
	fmt.Println("nextTime=", t)

	// Remove a specific job
	gocron.Remove(task)

	// Clear all scheduled jobs
	gocron.Clear()

	// Start all the pending jobs
	//<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(2).Seconds().Do(task)
	<-s.Start()
}
