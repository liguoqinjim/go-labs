package main

import (
	"lab092/lab002/data"
	"log"
)

func main() {
	//读取test.json
	data1, err := data.Asset("data/test.json")
	if err != nil {
		log.Fatalf("data.Asset error:%v", err)
	}
	log.Printf("data1=\n%s", data1)

	//读取hello.js
	data2, err := data.Asset("data/hello.js")
	if err != nil {
		log.Fatalf("data.Asset error:%v", err)
	}
	log.Printf("data2=\n%s", data2)
}
