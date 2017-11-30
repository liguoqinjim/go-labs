package main

import (
	"encoding/json"
	"github.com/xiocode/weigo"
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

	api := weigo.NewAPIClient(conf.AppKey, conf.AppSecret, conf.RedirectUrl, "code")
	api.SetAccessToken(conf.Token, 1519925461)

	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(weigo.Statuses)
	err := api.GET_statuses_home_timeline(kws, result)
	if err != nil {
		log.Println("err=", err)
	} else {
		log.Println(len(*result.Statuses))
	}
}

func ReadConf() {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	conf = &Conf{}
	json.Unmarshal(data, conf)

}
