package main

import (
	"fmt"
	"reflect"
)

func main() {
	//string to byte array
	a := "helloworld"
	fmt.Println("a.type=", reflect.TypeOf(a))
	b := []byte(a)
	fmt.Println("b.type=", reflect.TypeOf(b))
	fmt.Printf("%s\n", b)

	//byte array to string

	c := string(b[:])
	fmt.Println("c.type=", reflect.TypeOf(c))
	fmt.Println(c)
}
