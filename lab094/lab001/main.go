package main

import (
	"github.com/hugozhu/goweibo"
	"log"
	"os"
)

var sina = &weibo.Sina{
	AccessToken: weibo.ReadToken("./token"),
}

func main() {
	log.Println(os.Getenv("PWD"))

	//fetch 20 weibo after 12345678
	for _, p := range sina.TimeLine(0, "hugozhu", 12345678, 20) {
		log.Println(p)
	}
}
