package main

import (
	"log"
	"time"
)

func main() {
	tm := time.Now().Format(time.RFC1123)
	log.Println("RFC1123=", tm)
}
