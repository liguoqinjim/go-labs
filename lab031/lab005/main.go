package main

import (
	"fmt"
	"github.com/google/gops/agent"
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
	//gops
	if err := agent.Listen(nil); err != nil {
		fmt.Println(err)
	}

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("map.length=", len(Student_map))
			time.Sleep(time.Second * 4)
		}
	}()

	first_student_id := 1
	//缓慢增加缓存
	for i := 1; i < 12; i++ {
		endId := SimAddMemory(first_student_id)
		first_student_id = endId

		fmt.Println("first_student_id=", first_student_id)
		time.Sleep(time.Second * 3)
	}

	//为了gops来监控加的
	time.Sleep(time.Hour)
}
