package main

import (
	"lab092/lab002/data"
	"log"
)

func main() {
	data1,err := data.Asset("data/test.json")
	if err != nil{
		log.Println("Asset test.json error")
		return
	}
	log.Println("data1=",string(data1))

	data2,err := data.Asset("data/hello.js")
	if err != nil{
		log.Println("Asset hello.js error")
		return
	}
	log.Println("data2=",string(data2))
}