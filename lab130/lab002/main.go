package main

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
	"log"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func main() {
	list := arraylist.New()
	list.Add(&Student{Name: "tom", Age: 25, Score: 100})
	list.Add(&Student{Name: "Alice", Age: 14, Score: 90})
	list.Add(&Student{Name: "Ben", Age: 15, Score: 60})

	log.Println(list)

	//按age排序
	list.Sort(utils.Comparator(ageSort))
	log.Println("按age排序:")
	log.Println(list)

	//按score排序
	list.Sort(utils.Comparator(scoreSort))
	log.Println("按score排序:")
	log.Println(list)
}

func ageSort(a, b interface{}) int {
	s1 := a.(*Student)
	s2 := b.(*Student)

	if s1.Age < s2.Age {
		return -1
	} else if s1.Age > s2.Age {
		return 1
	} else {
		return 0
	}
}

func scoreSort(a, b interface{}) int {
	s1 := a.(*Student)
	s2 := b.(*Student)

	if s1.Score < s2.Score {
		return -1
	} else if s1.Score > s2.Score {
		return 1
	} else {
		return 0
	}
}
