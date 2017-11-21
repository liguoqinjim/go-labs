package main

import (
	"log"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/a/b/c", os.ModePerm) //生成多级目录
	if err != nil {
		log.Println(err)
	}
	log.Println("创建文件夹" + dir + "/a/b/c成功")
}
