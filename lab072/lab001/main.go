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
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
