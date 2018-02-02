package main

import (
	"github.com/mailru/easyjson"
	"lab121/lab001/data"
	"log"
)

func main() {
	//marshal
	m := &data.Message{Name: "Alice", Body: "Hello World", Time: 15000}
	err := easyjson.Marshal()
}
