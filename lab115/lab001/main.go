package main

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}

	md5h := md5.New()
	io.Copy(md5h, f)
	log.Printf("md5=%x\n", md5h.Sum([]byte("")))
}
