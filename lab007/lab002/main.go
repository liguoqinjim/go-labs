package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("os.Create error:%v", err)
	}
	defer f.Close()

	data := []byte("Hello\ngo\n")
	n, err := f.Write(data)
	if err != nil {
		log.Fatalf("f.Write error:%v", err)
	}
	log.Printf("wrote %d bytes", n)

	n2, err := f.WriteString("Hello Golang!")
	if err != nil {
		log.Fatalf("f.WriteString error:%v", err)
	}
	log.Printf("wrote %d bytes", n2)

	//sync
	if err := f.Sync(); err != nil {
		log.Fatalf("sync error:%v", err)
	}
}
