package main

import (
	"github.com/willf/bloom"
	"log"
)

func main() {
	n := uint(1000)
	filter := bloom.New(20*n, 5) // load of 20, 5 keys
	filter.Add([]byte("Love"))

	if filter.Test([]byte("Love")) {
		log.Println("already has")
	}

	if filter.TestString("hello") {
		log.Println("already has")
	} else {
		log.Println("not has")
	}
}
