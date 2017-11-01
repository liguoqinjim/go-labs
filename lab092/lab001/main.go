package main

import (
	"lab092/lab001/data"
	"log"
)

func main() {
	data, err := data.Asset("test.json")
	if err != nil {
		// Asset was not found.
	}

	// use asset data
	log.Println("data=", string(data))
}
