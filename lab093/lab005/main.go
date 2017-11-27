package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"errors"
)

//一旦找到就返回
var specDirName = "dir2" //要查找的文件夹
var foundPath = ""
var foundDir = errors.New("found directory")

func visit(path string, info os.FileInfo, err error) error {
	log.Printf("Visited: %s\n", path)

	if err != nil {
		log.Println("error:", err)
		return err
	}

	if info.IsDir() {
		pathNames := strings.Split(path, "\\")
		for _, v := range pathNames {
			if v == specDirName {
				foundPath = path
				return foundDir
			}
		}
	}

	return nil
}

//找出所有再返回

func main() {
	//根目录
	flag.Parse()
	root := flag.Arg(0)

	//一旦找到就返回
	err := filepath.Walk(root, visit)
	if err == foundDir {
		log.Println("找到文件夹")
	}
}
