package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get?a=1&b=2", nil)
	if err != nil {
		log.Fatalf("New Request error:%v", err)
	}
	//设置header
	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do error:%v", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("readAll error:%v", err)
	}
	log.Printf("Get返回:\n%s", string(b))
}
