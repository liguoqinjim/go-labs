package main

import "log"

type student struct {
	Name string
	Age  int
}

func main() {
	lab001()
	lab002()
	lab003()
}

func lab001() {
	m := make([]*student, 3)
	stus := []student{
		student{Name: "sa", Age: 10},
		student{Name: "sb", Age: 11},
		student{Name: "sc", Age: 12},
	}

	log.Println("################ 错误做法 ##################")
	for k, stu := range stus {
		m[k] = &stu
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}

	log.Println("################ 正确做法 ##################")
	for k, _ := range stus {
		m[k] = &stus[k]
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}
}

func lab002() {
	m := make([]student, 3)
	stus := []student{
		student{Name: "sa", Age: 10},
		student{Name: "sb", Age: 11},
		student{Name: "sc", Age: 12},
	}

	log.Println("################ 正确做法 ##################")
	for k, stu := range stus {
		m[k] = stu
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}

	log.Println("################ 正确做法 ##################")
	for k, _ := range stus {
		m[k] = stus[k]
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}
}

func lab003() {
	m := make([]*student, 3)
	stus := []*student{
		{Name: "sa", Age: 10},
		{Name: "sb", Age: 11},
		{Name: "sc", Age: 12},
	}

	log.Println("################ 正确做法 ##################")
	for k, stu := range stus {
		m[k] = stu
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}

	log.Println("################ 正确做法 ##################")
	for k, _ := range stus {
		m[k] = stus[k]
	}
	for _, s := range m {
		log.Println(s.Name, s.Age)
	}
}
