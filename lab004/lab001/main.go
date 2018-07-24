package main

import (
	"github.com/bitly/go-simplejson"
	"log"
)

func main() {
	//get content in json
	var testJson = []byte(`[
		{"Name": "Alice", "Age": 13},
		{"Name": "Bob",    "Age": 15}
	]`)

	json1, err := simplejson.NewJson(testJson)
	if err != nil {
		log.Fatalf("simplejson.NewJson error:%v", err)
	}

	json2, err := json1.Array()
	if err != nil {
		log.Fatalf("json1.Array error:%v", err)
	}

	for i := 0; i < len(json2); i++ {
		name, err := json1.GetIndex(i).Get("Name").String()
		if err != nil {
			log.Fatalf("Get string error:%v", err)
		}
		age, err := json1.GetIndex(i).Get("Age").Int()
		if err != nil {
			log.Fatalf("Get int error:%v", err)
		}
		log.Printf("name=%s,age=%d", name, age)
	}

	//marshal
	log.Println()
	js := simplejson.New()
	js.Set("Name", "Bruce")
	js.Set("Age", 20)
	b, err := js.Encode()
	if err != nil {
		log.Fatalf("js.Encode error:%v", err)
	}
	log.Printf("b = %s", b)
}
