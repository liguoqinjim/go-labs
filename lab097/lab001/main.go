package main

import (
	"github.com/otiai10/gosseract"
	"log"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("test.png")
	text, _ := client.Text()
	log.Println(text)
	// Hello, World!
}
