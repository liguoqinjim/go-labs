package main

import (
	"reflect"
	"log"
)

func main() {
	//string -> []byte
	a := "Hello World"
	log.Println("a.type=", reflect.TypeOf(a))
	b := []byte(a)
	log.Println("b.type=", reflect.TypeOf(b))
	log.Println("b=", b)

	//[]byte -> string
	c := string(b[:])
	log.Println("c.type=", reflect.TypeOf(c))
	log.Println("c=", c)
}
