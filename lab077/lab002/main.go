package main

import (
	"log"
	"sync"
)

var mu sync.Mutex

func main() {
	f()
}

func f() {
	mu.Lock()
	g()
	log.Println("call f()")

	mu.Unlock()
}

func g() {
	mu.Lock()
	log.Println("call g()")
	mu.Unlock()
}
