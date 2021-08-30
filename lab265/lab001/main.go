package main

import (
	"log"
	"runtime"
)

func main() {
	demo()
}

func demo() {
	log.Println(runtime.GOOS)
	log.Println(runtime.GOARCH)

	if runtime.GOOS == "darwin" {
		log.Println("is on mac")
	}
}
