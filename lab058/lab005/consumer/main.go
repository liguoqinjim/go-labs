package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	address := readConf()
	log.Println("address=", address)
}

func readConf() string {
	file, err := os.Open("ip.conf")
	if err != nil {
		log.Fatal("file open error:", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("readAll error:", err)
	}

	return string(data)
}
