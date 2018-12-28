package main

import (
	"log"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	//只是把Part1里面的字段换了一下顺序
	e byte
	c int8
	a bool
	b int32
	d int64
}

func main() {
	//golang中的一些数据类型的内存大小
	log.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
	log.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	log.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	log.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
	log.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
	log.Printf("string size: %d\n", unsafe.Sizeof("EDDYCJY"))

	//Part1的内存大小
	part1 := Part1{}
	log.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))

	//Part2的内存大小
	part2 := Part2{}
	log.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
}
