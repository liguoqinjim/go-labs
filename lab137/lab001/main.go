package main

import (
	"io/ioutil"
	"log"
)

func main() {
	confData, err := ioutil.ReadFile("etcd.conf")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error", err)
	}

	log.Println(string(confData))
}
