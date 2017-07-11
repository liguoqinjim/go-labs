package main

import (
	"log"
	"reflect"
	"strconv"
)

func main() {
	sa := "123.45"
	na1, err := strconv.ParseFloat(sa, 64)
	if err != nil {
		log.Fatalln("failed to parseFloat1:", err)
	}
	na2, err := strconv.ParseFloat(sa, 32)
	if err != nil {
		log.Fatalln("failed to parseFloat2:", err)
	}
	log.Println("64bitSize:", na1, reflect.TypeOf(na1))
	log.Println("32bitSize:", na2, reflect.TypeOf(na2))
}
