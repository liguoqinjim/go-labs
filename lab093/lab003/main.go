package main

import (
	"os"
	"path/filepath"
	"log"
)

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if f.IsDir() {
		log.Printf("visited dir: %s", path)
	} else {
		log.Printf("visited: %s", path)
	}
	return nil
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd error:%v", err)
	}
	log.Printf("pwd=%s", pwd)

	//遍历路径
	err = filepath.Walk(pwd, visit)
	if err != nil {
		log.Printf("filepath.Walk error:%v", err)
	}
}
