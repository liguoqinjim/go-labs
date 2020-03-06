package main

import (
	"github.com/mitchellh/go-homedir"
	"log"
)

func main() {
	//get home dir
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("homedir.Dir error:%v", err)
	}
	log.Println("dir:", dir)

	//把带~的路径转换为绝对路径
	expand, err := homedir.Expand("~/Workspace")
	if err != nil {
		log.Fatalf("homedir.Expand error:%v", err)
	}
	log.Println("expand:", expand)
}
