package main

import (
	"io/ioutil"
	"log"
)

type ConfStruct struct {
	Key string `json:"key"`
}

var Conf *ConfStruct

func main() {
	//读取配置文件
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}
	log.Printf("data=\n%s", data)

	//拼接url


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
