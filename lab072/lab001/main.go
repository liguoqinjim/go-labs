package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//float32范围
	fmt.Println(math.MaxFloat32)
	//float64范围
	fmt.Println(math.MaxFloat64)

	//float32占用内存
	var a float32 = 0.5
	var b float64 = 0.5
	var c int = 1
	fmt.Println("size of float32:", unsafe.Sizeof(a))
	fmt.Println("size of float64:", unsafe.Sizeof(b))
	fmt.Println("size of int:", unsafe.Sizeof(c))
}
