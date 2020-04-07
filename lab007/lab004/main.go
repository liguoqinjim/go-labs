package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("os.OpenFile error:%v", err)
	}
	defer f.Close()

	f.WriteString("test append\n")
	f.WriteString("test append\n")
}
