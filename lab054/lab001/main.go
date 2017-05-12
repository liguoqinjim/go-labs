package main

import (
	"github.com/tidwall/gjson"
)

func main() {
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value1 := gjson.Get(json, "name.last")
	println(value1.String())
	value2 := gjson.Get(json, "age")
	println(value2.Int())
}
