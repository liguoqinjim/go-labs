package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data := []byte("hello\ngo\n")

	err := ioutil.WriteFile("test.txt", data, 0644)
	if err != nil {
		log.Fatalf("writeFile error:%v", err)
	}
}
