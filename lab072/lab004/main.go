package main

import (
	"fmt"
	"reflect"
)

func main() {
	//float64
	var na float64 = 10.1
	var nb float64 = 10.2
	fmt.Println(na, nb)
	nc := na + nb
	fmt.Println(nc)
	fmt.Println(reflect.TypeOf(nc))

	//float32
	var nd float32 = 10.1
	var ne float32 = 10.2
	fmt.Println(nd, ne)
	nf := nd + ne
	fmt.Println(nf)
	fmt.Println(reflect.TypeOf(nf))
}
