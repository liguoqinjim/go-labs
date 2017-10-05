package main

import (
	"github.com/tidwall/gjson"
	"log"
)

var json = `{
	"type": "Dog",
	"Sound": "Bark",
	"Age": "11"
}`

var json2 = `
[{"type": "Dog",
	"Sound": "Bark",
	"Age": "11"},
{"type": "Cat",
	"Sound": "MiaoMiao",
	"Age": "3"}
]
`

func main() {
	m, ok := gjson.Parse(json).Value().(map[string]interface{})
	if !ok {
		// not a map
		log.Println("json not a map")
	} else {
		log.Println(m)
	}

	m, ok = gjson.Parse(json2).Value().(map[string]interface{})
	if !ok {
		log.Println("json2 not a map")
	} else {
		log.Println(m)
	}
}
