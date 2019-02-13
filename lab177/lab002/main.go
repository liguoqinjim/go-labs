package main

import (
	"log"
	"os"
)

func main() {
	//注意：newpath也要跟着文件名

	if ok, _ := exists("../data/3/"); !ok {
		log.Fatalf("路径不存在")
	}

	err := os.Rename("../data/1/1.txt", "../data/3/1.txt")
	if err != nil {
		panic(err)
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
