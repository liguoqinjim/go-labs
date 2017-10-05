package main

import (
	"github.com/tidwall/gjson"
	"log"
)

func main() {
	testJson := `
	{
	  "programmers": [
		{
		  "firstName": "Janet",
		  "lastName": "McLaughlin",
		}, {
		  "firstName": "Elliotte",
		  "lastName": "Hunter",
		}, {
		  "firstName": "Jason",
		  "lastName": "Harold",
		}
	  ]
	}
	`

	//这里的key会没有值
	value1 := gjson.Get(testJson, "programmers.#.lastName")
	log.Println("programmers.#.lastName=", value1)
	value1.ForEach(func(key, value gjson.Result) bool {
		log.Printf("value1:key=[%s],value=[%s]", key.String(), value.String())
		return true //keep iterating
	})

	//迭代对象的值
	value2 := gjson.Get(testJson, "programmers.0")
	log.Println("programmers.0=", value2)
	value2.ForEach(func(key, value gjson.Result) bool {
		log.Printf("value2:key=[%s],value=[%s]", key.String(), value.String())
		return true
	})
}
