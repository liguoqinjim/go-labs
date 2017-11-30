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

	//得到自己的微博
	kws := map[string]interface{}{
		"uid": "2684726573",
	}
	result := new(weigo.Statuses)
	err := api.GET_statuses_home_timeline(kws, result)
	if err != nil {
		log.Println("err=", err)
	} else {
		log.Println(len(*result.Statuses))
		for _, v := range *result.Statuses {
			log.Println(v.Mid)
		}
	}

	//发送微博
	kws = map[string]interface{}{
		"status": "Testing...Testing...",
	}
	result2 := new(weigo.Status)
	err = api.POST_statuses_update(kws, result2)
	if err != nil {
		log.Println("err=", err)
	} else {
		log.Println(result2)
	}

	//转发微博
	kws = map[string]interface{}{
		"id": "4179845682923414",
	}
	result3 := new(weigo.Status)
	err = api.POST_statuses_repost(kws, result3)
	if err != nil {
		log.Println("err=", err)
	} else {
		log.Println(result3)
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
