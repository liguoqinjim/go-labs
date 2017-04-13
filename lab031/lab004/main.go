package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

type Student struct {
	Sid     int
	Sname   string
	Sage    int
	Sscores [10]int
}

func NewStudent(id int) *Student {
	s := &Student{}
	s.Sid = id
	s.Sname = "xiaoming"
	s.Sage = 18
	s.Sscores = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return s
}

func SimAddMemory(startId int) (endId int) {
	endId = startId + 100000
	for i := startId; i <= endId; i++ {
		s := NewStudent(i)
		Student_map[i] = s
	}
	return
}

var Student_map = make(map[int]*Student)

func main() {
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

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("map.length=", len(Student_map))
			time.Sleep(time.Second * 4)
		}
	}()

	first_student_id := 1
	//缓慢增加缓存
	for i := 1; i < 10; i++ {
		endId := SimAddMemory(first_student_id)
		first_student_id = endId

		fmt.Println("first_student_id=", first_student_id)
		time.Sleep(time.Second * 5)
	}
}
