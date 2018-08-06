package main

import (
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("test.md")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	output := blackfriday.Run(data)
	err = ioutil.WriteFile("test.html", output, 0644)
	if err != nil {
		log.Fatalf("ioutil.WriteFile error:%v", err)
	}
}
