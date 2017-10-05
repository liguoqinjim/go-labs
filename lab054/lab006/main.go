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

	//存在
	value1 := gjson.Get(testJson, "programmers.0.firstName")
	if value1.Exists() {
		log.Println("programmers.0.firstName exists")
	} else {
		log.Println("programmers.0.firstName not exist")
	}

	//不存在
	value2 := gjson.Get(testJson, "programmers.firstName")
	if value2.Exists() {
		log.Println("programmers.firstName exists")
	} else {
		log.Println("programmers.firstName not exist")
	}
}
