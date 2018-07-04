package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
)

type ConfStruct struct {
	Key       string `json:"key"`
	EventName string `json:"event_name"`
	UrlFormat string `json:"url_format"`
	Value1    string `json:"value1"`
	Value2    string `json:"value2"`
	Value3    string `json:"value3"`
}

var conf = new(ConfStruct)

func main() {
	//读取配置文件
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}
	log.Printf("data=\n%s", data)

	err = json.Unmarshal(data, conf)
	if err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}

	//拼接url
	u := fmt.Sprintf(conf.UrlFormat, conf.EventName, conf.Key)
	log.Printf("url=%s", u)

	type ValueStruct struct {
		Value1 string `json:"value1"`
		Value2 string `json:"value2"`
		Value3 string `json:"value3"`
	}
	val := ValueStruct{Value1: conf.Value1, Value2: conf.Value2, Value3: conf.Value3}
	request := gorequest.New()
	//http://httpbin.org/post
	resp, body, errs := request.Post(u).
		Send(val).
		End()
	if errs != nil {
		log.Fatalf("errs:%+v", errs)
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll error:%v", err)
	}
	log.Printf("resp=\n%s", d)
	log.Println("body.length=", len(body))
	log.Println("body=", body)
}
