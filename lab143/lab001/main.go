package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ConfStruct struct {
	Key       string `json:"key"`
	EventName string `json:"event_name"`
	UrlFormat string `json:"url_format"`
	Value1    string `json:"value1"`
	Value2    string `json:"value2"`
	Value3    string `json:"value3"`
}

var conf *ConfStruct

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

	form := url.Values{
		"value1": {conf.Value1},
		"value2": {conf.Value2},
		"value3": {conf.Value3},
	}
	b := bytes.NewBufferString(form.Encode())
	resp, err := http.Post(u, "application/json", b)
	if err != nil {
		log.Fatalf("http.Post error:%v", err)
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll error:%v", err)
	}
	log.Printf("resp:\n%s", d)

	//在上面的post基础上指定contentType和body
	//form := url.Values{
	//	"username": {"xiaoming"},
	//	"address":  {"beijing"},
	//	"subject":  {"Hello"},
	//	"from":     {"china"},
	//}
	//postBody := bytes.NewBufferString(form.Encode())
	//resp2, err2 := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", postBody)
	//if err2 != nil {
	//	log.Fatalf("post error:%v", err2)
	//}
	//body2, err2 := ioutil.ReadAll(resp2.Body)
	//if err2 != nil {
	//	log.Fatalf("readAll error:%v", err2)
	//}
	//resp2.Body.Close()
	//log.Printf("post2\n%s", string(body2))
}
