package main

import (
	"github.com/jmz331/gpinyin"
	"log"
)

func main() {
	//简->繁
	const s = "国家体育场"
	r := gpinyin.ConvertToTraditionalChinese(s)
	log.Println(r)

	//繁->简
	const a = "國家體育場"
	r2 := gpinyin.ConvertToSimplifiedChinese(a)
	log.Println(r2)
}
