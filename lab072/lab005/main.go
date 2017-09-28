package main

import (
	"log"
	"math"
)

func main() {
	a := 123.45
	log.Println("a=", a)

	b := math.Nextafter(a, 123.46)
	log.Println("b=", b)

	c := math.Nextafter(a, 123.44)
	log.Println("c=", c)
}
