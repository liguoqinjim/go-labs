package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}
	fmt.Println("s=", s)
	fmt.Printf("&s=%p\n", &s)

	p := unsafe.Pointer(&s)
	fmt.Println("p=", p)
	up0 := uintptr(p)
	fmt.Println("up0=", up0)
	pb := (*byte)(p)
	*pb = 10
	fmt.Println(s)

	//偏移到第二个字段
	up := up0 + unsafe.Offsetof(s.b)
	fmt.Println("up=", up)
	p = unsafe.Pointer(up)
	pb = (*byte)(p)
	*pb = 20
	fmt.Println(s)

	//偏移到第三个字段
	up = up0 + unsafe.Offsetof(s.c)
	fmt.Println("up=", up)
	p = unsafe.Pointer(up)
	pb = (*byte)(p)
	*pb = 50
	fmt.Println(s)

	//偏移到第四个字段
	up = up0 + unsafe.Offsetof(s.d)
	fmt.Println("up=", up)
	p = unsafe.Pointer(up)
	pi := (*int64)(p)
	*pi = 40
	fmt.Println(s)
}
