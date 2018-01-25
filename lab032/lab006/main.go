package main

import (
	"lab032/lab006/consts"
	"lab032/lab006/utils"
	"log"
)

func main() {
	log.Println("hello main,", consts.HELLO_MESSAGE)

	c := utils.Add(5, 6)
	log.Println("c=", c)
}
