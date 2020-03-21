package main

import (
	"log"
	"os"
)

func main() {
	//os.Remove不可以删除
	//if err := os.Remove("data/1"); err != nil {
	//	log.Fatalf("os.Remove error:%v", err)
	//}

	if err := os.RemoveAll("data/1"); err != nil {
		log.Fatalf("os.RemoveAll error:%v", err)
	}
}
