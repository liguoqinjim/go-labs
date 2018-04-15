package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}

	form := url.Values{
		"username": {"xiaoming"},
		"address":  {"beijing"},
		"subject":  {"Hello"},
		"from":     {"china"},
	}
	postBody := bytes.NewBufferString(form.Encode())

	req, err := http.NewRequest(http.MethodPost, "http://httpbin.org/post", postBody)
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
	log.Printf("Post1返回:\n%s", string(b))

	client2 := &http.Client{}
	form2 := url.Values{}
	form2.Set("username", "gao")
	postBody2 := strings.NewReader(form2.Encode())
	req2, err2 := http.NewRequest(http.MethodPost, "http://httpbin.org/post", postBody2)
	if err2 != nil {
		log.Fatalf("NewRequest error:%v", err2)
	}
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp2, err2 := client2.Do(req2)
	if err2 != nil {
		log.Fatalf("client do error:%v", err2)
	}
	b2, err2 := ioutil.ReadAll(resp2.Body)
	if err2 != nil {
		log.Fatalf("readAll error:%v", err2)
	}
	log.Printf("Post2返回:\n%s", string(b2))
}
