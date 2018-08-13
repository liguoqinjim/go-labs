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
var specDirName = "liguoqinjim" //要查找的文件夹

//一旦找到就返回
var foundDir = errors.New("found directory")

func main() {
	//一旦找到就返回
	path, found := findDir(root, specDirName, -1)
	if found {
		log.Printf("found,path=%s", path)
	} else {
		log.Println("not found")
	}

	//找出所有
	paths, found := findDirAll(root, specDirName, -1)
	if found {
		log.Printf("foundAll,paths=%v", paths)
	} else {
		log.Println("not foundAll")
	}
}

//寻找文件夹，找到一个就退出，depth可以指定寻找的文件夹深度，-1的时候不会考虑深度
func findDir(root, dir string, depth int) (string, bool) {
	foundPath := ""
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//判断深度
		if depth > 0 {
			rootCount := strings.Count(root, "\\")
			if rootCount == 0 {
				rootCount = strings.Count(root, "/")
			}

			dep := strings.Count(path, string(os.PathSeparator)) - rootCount
			if dep > depth {
				return filepath.SkipDir
			}
		}

		if info.IsDir() {
			pathNames := strings.Split(path, string(os.PathSeparator))
			for _, v := range pathNames {
				if v == dir {
					foundPath = path
					return foundDir
				}
			}
		}

		return nil
	})

	if err == foundDir {
		return foundPath, true
	} else {
		return foundPath, false
	}
}

func findDirAll(root, dir string, depth int) ([]string, bool) {
	var foundpaths []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("error:", err)
			return err
		}

		//判断深度
		if depth > 0 {
			rootCount := strings.Count(root, "\\")
			if rootCount == 0 {
				rootCount = strings.Count(root, "/")
			}
			dep := strings.Count(path, string(os.PathSeparator)) - rootCount
			if dep > depth {
				return filepath.SkipDir
			}
		}

		//判断是否是文件夹
		if info.IsDir() {
			pathNames := strings.Split(path, string(os.PathSeparator))
			for _, v := range pathNames {
				if v == dir {
					foundpaths = append(foundpaths, path)
				}
			}
		}

		return nil
	})

	if err == foundDir {
		return foundpaths, true
	} else {
		if len(foundpaths) > 0 {
			return foundpaths, true
		} else {
			return foundpaths, false
		}
	}
}
