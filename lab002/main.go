package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	//marshal object->json
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("b = %s\n", b)

	//unmarshal json->object
	b2 := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var m2 Message
	err = json.Unmarshal(b2, &m2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("m2 = %+v\n", m2)
}
