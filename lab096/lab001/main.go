package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Conf struct {
	AppKey      string `json:"AppKey"`
	AppSecret   string `json:"AppSecret"`
	RedirectUrl string `json:"RedirectUrl"`
	Token       string `json:"Token"`
}

var conf *Conf

func main() {
	ReadConf()
}

func ReadConf() {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	conf = &Conf{}
	json.Unmarshal(data, conf)

}
