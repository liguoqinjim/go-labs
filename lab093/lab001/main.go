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
		log.Println(v.Name())
	}
}
