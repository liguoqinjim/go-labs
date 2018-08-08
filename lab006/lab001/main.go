package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}
	log.Printf("data = %s\n", data)
}
