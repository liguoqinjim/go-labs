package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("os.Create error:%v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	n, err := w.WriteString("Hello go!\n")
	if err != nil {
		log.Fatalf("WriteString error:%v", err)
	}
	log.Printf("wrote %d bytes\n", n)

	if err := w.Flush(); err != nil {
		log.Fatalf("w.Flush error")
	}
}
