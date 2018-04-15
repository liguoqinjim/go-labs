package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//最简单的post
	resp1, err1 := http.Post("http://httpbin.org/post", "", nil)
	if err1 != nil {
		log.Fatalf("post error:%v", err1)
	}
	body1, err1 := ioutil.ReadAll(resp1.Body)
	if err1 != nil {
		log.Fatalf("readAll error:%v", err1)
	}
	resp1.Body.Close()
	log.Printf("post1\n%s", string(body1))

	//在上面的post基础上指定contentType和body
	form := url.Values{
		"username": {"xiaoming"},
		"address":  {"beijing"},
		"subject":  {"Hello"},
		"from":     {"china"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	resp2, err2 := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", postBody)
	if err2 != nil {
		log.Fatalf("post error:%v", err2)
	}
	body2, err2 := ioutil.ReadAll(resp2.Body)
	if err2 != nil {
		log.Fatalf("readAll error:%v", err2)
	}
	resp2.Body.Close()
	log.Printf("post2\n%s", string(body2))

	//postForm方法
	resp, err := http.PostForm("http://httpbin.org/post",
		url.Values{"wd": {"github"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("post3\n%s", string(body))
}
