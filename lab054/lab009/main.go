package main

import (
	"github.com/tidwall/gjson"
	"log"
)

type Animals struct {
	Animals []*Animal `json:"animals"`
}

type Animal struct {
	Type  string `json:"type"`
	Sound string `json:"sound"`
	Age   int    `json:"age"`
}

var json = `
{"animals":[{"type": "Dog",
	"Sound": "Bark",
	"Age": "11"},
{"type": "Cat",
	"Sound": "MiaoMiao",
	"Age": "3"}
]}
`

func main() {
	var animals Animals
	gjson.Unmarshal([]byte(json), &animals)
	log.Printf("animals=%+v", animals)
	for n, v := range animals.Animals {
		log.Printf("animals[%d]=%+v", n, v)
	}
}
