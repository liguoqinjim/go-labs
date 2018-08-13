package main

import (
	"os"
	"log"
)

func main() {
	exist, err := exists("E:/Workspace/go/src/github.com")
	if exist {
		log.Println("exists")
	} else {
		log.Println("not exists")
		if err != nil {
			log.Fatalf("error:%v", err)
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
