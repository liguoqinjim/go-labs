package main

import (
	"encoding/json"
	"log"
)

func main() {
	js1 := `{"hello":"world"}`
	js2 := `"hello":"world"`
	js3 := `{"hello"}`
	js4 := `111`
	js5 := `[{"key": "value1"}, {"key": "value2"}]`

	//js1
	if json.Valid([]byte(js1)) {
		log.Println("js1 is valid")
	} else {
		log.Println("js1 is not valid")
	}

	//js2
	if json.Valid([]byte(js2)) {
		log.Println("js2 is valid")
	} else {
		log.Println("js2 is not valid")
	}

	//js3
	if json.Valid([]byte(js3)) {
		log.Println("js3 is valid")
	} else {
		log.Println("js3 is not valid")
	}

	//js4
	if json.Valid([]byte(js4)) {
		log.Println("js4 is valid")
	} else {
		log.Println("js4 is not valid")
	}

	//js5
	if json.Valid([]byte(js5)) {
		log.Println("js5 is valid")
	} else {
		log.Println("js5 is not valid")
	}
}
