package main

import (
	"io/ioutil"
	"log"
)

func main() {
	dir, e := ioutil.ReadDir("E:/Workspace/SLGDev/SLGServer")
	if e != nil {
		log.Println("read dir error")
		return
	}
	for _, v := range dir {
		if v.IsDir() {
			log.Printf("文件夹:%s", v.Name())
		} else {
			log.Println(v.Name())
		}
	}
}
