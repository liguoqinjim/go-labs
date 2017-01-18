package main

import (
	"fmt"
	"reflect"
)

type S struct {
	F string `species:"gopher" color:"blue"`
}

type S2 struct {
	F0 string `alias:"field_0"`
	F1 string `alias:""`
	F2 string
}

func main() {
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)

	color := field.Tag.Get("color")
	fmt.Println("color =", color)

	species := field.Tag.Get("species")
	fmt.Println("species =", species)

	//返回空
	test := field.Tag.Get("test")
	fmt.Println("test =", test)
}
