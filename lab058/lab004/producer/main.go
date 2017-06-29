package main

import (
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	config := nsq.NewConfig()
	_ = config

	address := readConf()
	log.Println("address=", address)
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
