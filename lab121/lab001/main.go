package main

import (
	"lab121/lab001/data"
	"github.com/mailru/easyjson"
	"log"
)

func main() {
	//marshal
	m := &data.Message{Name:"Alice",Body:"Hello World",Time:15000}
	err := easyjson.Marshal()
	}
