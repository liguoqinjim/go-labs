package main

import (
	"os"
	"runtime/trace"
)

func fakeReadDB() {
	ms := make([]int, 100)
	for n := range ms {
		ms[n] = n
	}

	total := 0
	for _, v := range ms {
		total += v
	}
}

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

	// Your program here
	for i := 0; i < 1000; i++ {
		fakeReadDB()
	}
}
