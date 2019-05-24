package main

import (
	"log"
	"sync"
)

func main() {
	var a1 sync.Map
	var b1 sync.Map

	for i := 0; i < 10; i++ {
		b1.Store(i, i+10)
	}

	//注意，是b1的指针
	a1.Store("b1", &b1)
	b1.Range(func(key, value interface{}) bool {
		log.Println(key, value)
		return true
	})
	b1.Store("aaa", "qqq")

	log.Println("a1--------------")
	if m, ok := a1.Load("b1"); ok {
		mm := m.(*sync.Map)

		mm.Range(func(key, value interface{}) bool {
			log.Println(key, value)
			return true
		})
	}
}
