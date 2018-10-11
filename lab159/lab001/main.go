package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println(rand.Intn(100))

	log.Println(rand.Float64())

	s1 := rand.NewSource(time.Now().UnixNano()) //创建随机数种子
	r1 := rand.New(s1)

	log.Println(r1.Intn(100))

	r2 := rand.New(rand.NewSource(2))
	r3 := rand.New(rand.NewSource(2))
	log.Println(r2.Intn(100), r2.Intn(100))
	log.Println(r3.Intn(100), r3.Intn(100))
}
