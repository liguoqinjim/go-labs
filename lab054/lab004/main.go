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

	//只返回匹配到的第一个值
	value1 := gjson.Get(testJson, `friends.#[last=="Murphy"].first`)
	log.Println(`friends.#[last="Murphy"].first=`, value1.String())

	//返回所有匹配到的值，区别是所有匹配的时候要用两个#号
	value2 := gjson.Get(testJson, `friends.#[last="Murphy"]#.first`)
	log.Println(`friends.#[last="Murphy"]#.first=`, value2)
	if value2.IsArray() {
		for n, v := range value2.Array() {
			log.Printf("value2[%d]=%s", n, v.String())
		}
	}

	value3 := gjson.Get(testJson, `friends.#[age>45]#.first`)
	log.Println(`friends.#[age>45]#.first`, value3)
	if value3.IsArray() {
		for n, v := range value3.Array() {
			log.Printf("value3[%d]=%s", n, v.String())
		}
	}

	//注意：就算结果只有一个，但是返回的还是数组
	value4 := gjson.Get(testJson, `friends.#[age>47]#.first`)
	log.Println(`friends.#[age>47]#.first`, value4)
	if value4.IsArray() {
		for n, v := range value4.Array() {
			log.Printf("value4[%d]=%s", n, v.String())
		}
	}

	value5 := gjson.Get(testJson, `friends.#[first%"D*"].first`)
	log.Println(`friends.#[first%"D*"].first`, value5)
}
