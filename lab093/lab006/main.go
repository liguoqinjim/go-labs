package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"errors"
)

//根路径
var root = "E:/Workspace/go/src/github.com"

//var root = "dir1"
var specDirName = "liguoqinjim" //要查找的文件夹

//一旦找到就返回
var foundPath = ""
var foundDir = errors.New("found directory")

func visit(path string, info os.FileInfo, err error) error {
	log.Printf("Visited: %s\n", path)

	if err != nil {
		log.Println("error:", err)
		return err
	}

	//判断深度
	rootCount := strings.Count(root, "\\")
	if rootCount == 0 {
		rootCount = strings.Count(root, "/")
	}
	depth := strings.Count(path, "\\") - rootCount
	//log.Printf("depth[%d] [%d,%d]", depth, strings.Count(path, "\\"), strings.Count(root, "\\"))
	if depth > specDepth {
		return filepath.SkipDir
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
var paths []string
var specDepth = 1

func visitAll(path string, info os.FileInfo, err error) error {
	log.Printf("Visited: %s\n", path)

	if err != nil {
		log.Println("error:", err)
		return err
	}

	//判断深度
	rootCount := strings.Count(root, "\\")
	if rootCount == 0 {
		rootCount = strings.Count(root, "/")
	}
	depth := strings.Count(path, "\\") - rootCount
	//log.Printf("depth[%d] [%d,%d]", depth, strings.Count(path, "\\"), strings.Count(root, "\\"))
	if depth > specDepth {
		return filepath.SkipDir
	}

	//判断是否是文件夹
	if info.IsDir() {
		pathNames := strings.Split(path, "\\")
		for _, v := range pathNames {
			if v == specDirName {
				paths = append(paths, path)
			}
		}
	}

	return nil
}

func main() {
	if root == "" {
		log.Fatalf("root error:%s", root)
	}

	//一旦找到就返回
	log.Println("方式1-------------------------------------------------------------")
	err := filepath.Walk(root, visit)
	if err == foundDir {
		log.Println("找到文件夹")
	} else {
		log.Println("没有找到")
	}

	//找出所有
	log.Println("方式2-------------------------------------------------------------")
	err = filepath.Walk(root, visitAll)
	if paths == nil {
		log.Println("没有找到")
	} else {
		log.Println("找到的文件夹", paths)
	}
}
