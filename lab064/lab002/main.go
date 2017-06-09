package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ConnectInfo struct {
	Username string
	Pwd      string
	Hostname string
	Port     string
	DB       string
}

func readConf() *ConnectInfo {
	file, err := os.Open("mongo.json")
	handleError(err)

	data, err := ioutil.ReadAll(file)
	handleError(err)

	ci := new(ConnectInfo)
	err = json.Unmarshal(data, ci)
	handleError(err)

	return ci
}

var connectInfo *ConnectInfo

func main() {
	connectInfo = readConf()
	fmt.Println(connectInfo)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
