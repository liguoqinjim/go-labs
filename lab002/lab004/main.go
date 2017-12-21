package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatalf("marshal error:%v", err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	log.Println(out.String())
}
