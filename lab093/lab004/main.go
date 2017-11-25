package main

import (
	"flag"
	"github.com/henrylee2cn/faygo/errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var filterPath = "dir2"

func visit(path string, f os.FileInfo, err error) error {
	log.Println("Visited:", path, "--------------------------------------------------------")
	log.Println("是否是文件夹", f.IsDir())

	if strings.Contains(path, filterPath) {
		return errors.New(path)
	} else {
		return nil
	}
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)

	if err == nil {
		log.Println("没有找到文件夹")
	} else {
		log.Println("找到文件夹:", err.Error())
	}
}
