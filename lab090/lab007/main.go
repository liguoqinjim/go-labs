package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/url"
)

const (
	url1 = "http://httpbin.org/get"
	url2 = "http://httpbin.org/get?a=1&b=2"
)

func main() {
	request := gorequest.New()

	_, body, errs := request.Get(url1).
		Param("a", "1").
		Param("b", "高松").
		Param("c", url.QueryEscape("高松")).
		End()
	if errs != nil {
		log.Fatalf("request.Get errors:%v", errs)
	}
	log.Println("body=", body)
}
