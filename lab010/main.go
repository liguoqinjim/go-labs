package main

import (
	"io/ioutil"
	"log"
)

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("test.txt", d1, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
