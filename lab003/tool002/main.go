package main

import (
	"encoding/json"
	"log"
)

func main() {
	data := []byte(`{"Depth":0,"Ctx":{"page":1,"a":"b"}}`)

	r := new(serializableRequest)
	if err := json.Unmarshal(data, r); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	log.Println(r.Ctx["a"].(string))
	log.Println(r.Ctx["page"].(int))
}

type serializableRequest struct {
	Depth int
	Ctx   map[string]interface{}
}
