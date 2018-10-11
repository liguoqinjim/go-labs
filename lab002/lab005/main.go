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
	log.Println("data=", string(data))

	//解析第一层
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		log.Printf("json unmarshal error:%v", err)
	}
	log.Println("解析第一层:", objmap)

	//解析sendMsg：解析到结构体
	log.Println("sendMsg内容:", string(*objmap["sendMsg"]))
	var s sendMsg
	err = json.Unmarshal(*objmap["sendMsg"], &s)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Printf("s=%#v", s)

	//解析sendMsg：解析到map[string]*json.RawMessage
	var objmap3 map[string]*json.RawMessage
	err = json.Unmarshal(*objmap["sendMsg"], &objmap3)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println("objmap3=", objmap3)

	//解析str：解析到string
	var str string
	err = json.Unmarshal(*objmap["say"], &str)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println("str=", str)
	//解析str:解析到map[string]*json.RawMessage
	var objmap2 map[string]*json.RawMessage
	err = json.Unmarshal(*objmap["say"], &objmap2)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
	log.Println(objmap2)
}
