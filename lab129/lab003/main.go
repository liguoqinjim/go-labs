package main

import (
	"log"
	"sync"
)

type Person struct {
	Id   int
	Name string
}

var ps sync.Map

func main() {
	p1 := &Person{1, "tom"}
	p2 := &Person{2, "Kimi"}
	p3 := &Person{3, "Alice"}

	ps.Store(1, p1)
	ps.Store(2, p2)
	ps.Store(3, p3)

	ps.Range(func(key, value interface{}) bool {
		log.Println("key=", key)
		p := value.(*Person)
		log.Println("p=", p)
		p.Name += "_"

		return true
	})

	ps.Range(func(key, value interface{}) bool {
		p := value.(*Person)
		log.Println("p=", p)

		return true
	})
}
