package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://httpbin.org/get?a=1&b=2")
	if err != nil {
		log.Fatalf("http.Get error:%v", err)
	}
	defer resp.Body.Close()
	//读取body数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("readAll error:%v", err)
	}
	log.Printf("\n%s", string(body))
}
