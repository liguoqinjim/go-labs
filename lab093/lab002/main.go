package main

import (
	"log"
	"os"
)

func main() {
	//选择路径分割符，windows和linux不一样
	var pathSeparator string
	if os.IsPathSeparator('\\') { //windows
		pathSeparator = "\\"
	} else {
		pathSeparator = "/"
	}
	log.Println("pathSeparator=", pathSeparator)

	//当前路径
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd error:%v", err)
	}
	log.Println("pwd=", pwd)

	//创建文件夹
	newDir := pwd + pathSeparator + "md"
	err = os.Mkdir(newDir, os.ModePerm) //在当前目录下生成md目录
	if err != nil {
		log.Fatalf("os.Mkdir error:%v", err)
	}
	log.Printf("mkdir:%s", newDir)

	//创建多级目录
	newDir2 := pwd + pathSeparator + "a" + pathSeparator + "b"
	err = os.MkdirAll(newDir2, os.ModePerm) //生成多级目录
	if err != nil {
		log.Fatalf("os.MkdirAll error:%v", err)
	}
	log.Printf("mkdirAll:%s", newDir2)
}
