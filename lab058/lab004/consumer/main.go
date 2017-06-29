package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	ip := readConf()
	log.Println("ip=", ip)
}

func readConf() string {
	file, err := os.Open("ip.conf")
	if err != nil {
		log.Fatal("open file error:", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("readAll error:", err)
	}

	return string(data)
}
