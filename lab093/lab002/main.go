package main

import (
	"log"
	"os"
)

func main() {
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	log.Println(path)

	dir, _ := os.Getwd()                        //当前的目录
	err := os.Mkdir(dir+path+"md", os.ModePerm) //在当前目录下生成md目录
	if err != nil {
		log.Println(err)
	}
	log.Println("创建目录" + dir + path + "md成功")
}
