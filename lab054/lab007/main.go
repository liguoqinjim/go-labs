package main

import (
	"github.com/tidwall/gjson"
	"log"
)

type Animal struct {
	Type  string `json:"type"`
	Sound string `json:"sound"`
	Age   int    `json:"age"`
}

var json = `{
	"type": "Dog",
	"Sound": "Bark",
	"Age": "11"
}`

func main() {
	var dog Animal
	gjson.Unmarshal([]byte(json), &dog)
	log.Printf("dog=%+v", dog)
}
