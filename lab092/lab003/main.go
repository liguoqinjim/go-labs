package main

import (
	"lab092/lab003/assets1"
	"lab092/lab003/assets2"
	"log"
)

func main() {
	//读取static1里面的文件
	data1, err := assets1.Asset("static1/hello.js")
	if err != nil {
		log.Fatalf("assets1.Asset error:%v", err)
	}
	log.Printf("data1=\n%s", data1)

	//读取static2里面的文件
	data2, err := assets2.Asset("data.json")
	if err != nil {
		log.Fatalf("assets2.Asset error:%v", err)
	}
	log.Printf("data2=\n%s", data2)
}
