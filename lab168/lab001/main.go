package main

import (
	"github.com/spaolacci/murmur3"
	"log"
)

func main() {
	h32 := murmur3.New32WithSeed(0)
	h32.Write([]byte("hello"))
	v := h32.Sum32()
	log.Println(v)
}
