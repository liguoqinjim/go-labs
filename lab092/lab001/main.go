package main

import (
	"lab092/lab001/data"
	"log"
)

func main() {
	data, err := data.Asset("test.json")
	if err != nil {
		log.Fatalf("data.Asset error:%v", err)
	}

	// use asset data
	log.Printf("data=\n%s", data)
}
