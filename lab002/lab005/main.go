package main

import (
	"encoding/json"
	"log"
)

// the struct for the value of a "sendMsg"-command
type sendMsg struct {
	User string
	Msg  string
}

// The type for the value of a "say"-command
type say string

func main() {
	data := []byte(`{"sendMsg":{"user":"ANisus","msg":"Trying to send a message"},"say":"Hello"}`)

	//解析第一层
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		log.Printf("json unmarshal error:%v", err)
	}
	log.Println(objmap)

	//sendMsg
	log.Println(string(*objmap["sendMsg"]))
	var s sendMsg
	err = json.Unmarshal(*objmap["sendMsg"], &s)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println(s)
	//解析sendMsg
	var objmap3 map[string]*json.RawMessage
	err = json.Unmarshal(*objmap["sendMsg"], &objmap3)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println(objmap3)

	//str
	var str string
	err = json.Unmarshal(*objmap["say"], &str)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println(str)
	var objmap2 map[string]*json.RawMessage
	err = json.Unmarshal(*objmap["say"], &objmap2)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println(objmap2)
}
