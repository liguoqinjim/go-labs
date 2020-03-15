package main

import (
	"log"
	"os"
)

func main() {
	//os.Getwd
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pwd)
}
