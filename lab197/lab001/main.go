package main

import (
	"github.com/segmentio/ksuid"
	"log"
)

func main() {
	ks := ksuid.New()
	log.Println(ks.String())
	log.Println(ks.Next().String())
	log.Println(ks.Prev().String())

	for i := 0; i < 10; i++ {
		log.Println(ksuid.New().Next().String())
	}

	for i := 0; i < 10; i++ {
		log.Println(ksuid.New().String())
	}
}
