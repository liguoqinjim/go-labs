package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data = %s\n", data)
}
