package main

import (
	"io/ioutil"
	"log"
)

func main() {
	dir, err := ioutil.ReadDir("E:/Workspace/SLGDev/SLGServer")
	if err != nil {
		log.Fatalf("readDir error:%v", err)
	}

	for _, v := range dir {
		if v.IsDir() {
			log.Printf("文件夹:%s", v.Name())
		} else {
			log.Printf("文件:%s", v.Name())
		}
	}
}
