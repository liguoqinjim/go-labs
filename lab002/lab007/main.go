package main

import (
	"encoding/json"
	"log"
)

type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`

	Data  json.RawMessage `json:"data"`
	Data2 interface{}     `json:"data_2"`
}

func main() {
	var data = `{"a":1,"b":2,"c":3}`

	s := &Student{
		Name: "tom", Id: 1,
	}
	s.Data = []byte(data)
	s.Data2 = []byte(data)

	m, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("json.Marshal error:%v", err)
	}
	log.Printf("%s", m)
	//output:
	//{"name":"tom","id":1,"data":{"a":1,"b":2,"c":3},"data_2":"eyJhIjoxLCJiIjoyLCJjIjozfQ=="}
}
