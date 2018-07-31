package main

import (
	"regexp"
	"log"
)

func main() {
	//判断是否匹配正则表达式
	matched, err := regexp.MatchString("p([a-z]+)ch", "peach")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(matched)

	//判断是否匹配正则表达式
	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r.MatchString("peach"))

	//找出第一个匹配的值
	log.Println(r.FindString("peach punch"))

	//
	log.Println(r.FindStringSubmatch("peach punch"))

	//找出所有匹配的值
	log.Println(r.FindAllString("peach punch pinch", -1))
}
