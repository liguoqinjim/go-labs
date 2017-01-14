package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//NewDecoder
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock.",
		"Time":1}
		{"Name": "Sam", "Text": "Who's there?","Time":2}
		{"Name": "Ed", "Text": "Go fmt.","Time":3}
		{"Name": "Sam", "Text": "Go fmt who?","Time":4}
		{"Name": "Ed", "Text": "Go fmt yourself!","Time":5}
	`
	type Message struct {
		Name, Text string
		Time       float64
	}
	ms := make([]interface{}, 0)
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", m)
		ms = append(ms, m)
	}

	//NewEncoder
	fmt.Println()
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&ms); err != nil {
		fmt.Println(err)
	}
}
