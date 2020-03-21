package main

import (
	"github.com/willf/bloom"
	"log"
)

func main() {
	n := uint(1000)
	filter := bloom.NewWithEstimates(20*n, 5) // load of 20, 5 keys

	log.Println(filter.EstimateFalsePositiveRate(n))

	filter.Add([]byte("Love"))
	if filter.Test([]byte("Love")) {
		log.Println("Love already has")
	}

	if filter.TestString("Hello") {
		log.Println("Hello already has")
	} else {
		log.Println("Hello not has")
	}
}
