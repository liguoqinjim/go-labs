package main

import (
	"github.com/atotto/clipboard"
	"log"
)

func main() {
	clipboard.WriteAll("日本語123")
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalf("clipboard.ReadAll error:%v", err)
	}
	log.Println(text)
}
