package main

import (
	"lab120/lab001/data"
	"log"
)

func main() {
	//marshal
	m := data.Message{Name: "Alice", Body: "Hello World", Time: 15000}
	content, err := m.MarshalJSON()
	if err != nil {
		log.Printf("m.MarshalJSON error:%v", err)
	}
	log.Printf("content=%s", content)

	//unmarshal
	m2 := &data.Message{}
	err = m2.UnmarshalJSON(content)
	if err != nil {
		log.Printf("m2.UnmarshalJSON error:%v", err)
	}
	log.Println("m2=", m2)
}
