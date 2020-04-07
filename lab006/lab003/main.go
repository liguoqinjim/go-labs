package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer f.Close()

	//按行读取
	buf := bufio.NewReaderSize(f, 0)
	for {
		line, err := buf.ReadBytes('\n')
		log.Printf("%s", line)

		if err == io.EOF {
			break
		}
	}
}
