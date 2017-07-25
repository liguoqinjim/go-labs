package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	f()
}

func f() {
	mu.Lock()
	g()
	fmt.Println("call f()")

	mu.Unlock()
}

func g() {
	mu.Lock()
	fmt.Println("call g()")
	mu.Unlock()
}
