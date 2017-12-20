package main

import (
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("b = %s\n", b)
}
