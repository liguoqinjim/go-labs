package main

import (
	"github.com/tidwall/gjson"
	"log"
)

func main() {
	testJson := `
	{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "age":37,
	  "children": ["Sara","Alex","Jack"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "Dale", "last": "Murphy", "age": 44},
		{"first": "Roger", "last": "Craig", "age": 68},
		{"first": "Jane", "last": "Murphy", "age": 47}
	  ]
	}
	`

	value1 := gjson.Get(testJson, "name.last")
	log.Println("name.last=", value1.String())

	value2 := gjson.Get(testJson, "age")
	log.Println("age=", value2.Int())

	value3 := gjson.Get(testJson, "children")
	log.Println("children=", value3)
	if value3.IsArray() {
		for n, v := range value3.Array() {
			log.Printf("value3[%d]=%s", n, v.String())
		}
	}

	value4 := gjson.Get(testJson, "children.#")
	log.Println("children.#=", value4.Int())

	//index是从0开始
	value5 := gjson.Get(testJson, "children.1")
	log.Println("children.1=", value5.String())

	value6 := gjson.Get(testJson, "child*.2")
	log.Println("child*.2=", value6.String())

	value7 := gjson.Get(testJson, "c?ildren.0")
	log.Println("c?ildren.0=", value7.String())

	value8 := gjson.Get(testJson, "fav\\.movie")
	log.Println("fav\\.movie=", value8.String())

	value9 := gjson.Get(testJson, "friends.#.first")
	log.Println("friends.#.first=", value9)
	if value9.IsArray() {
		for n, v := range value9.Array() {
			log.Printf("value9[%d]=%s", n, v.String())
		}
	}

	value10 := gjson.Get(testJson, "friends.1.last")
	log.Println("friends.1.last=", value10.String())
}
