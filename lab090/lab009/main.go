package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
)

const (
	HR = "------------------------------------------------------------------"
)

func main() {
	//Post
	log.Println("Post" + HR)
	request := gorequest.New()
	resp, body, errs := request.Post("http://httpbin.org/post").End()
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(body))
	log.Println("body=", body)

	//JSON Post 01
	log.Println("Json Post 01" + HR)
	request = gorequest.New()
	resp, body, errs = request.Post("http://httpbin.org/post").
		Set("Notes", "gorequst is coming!").
		Send(`{"name":"backy", "species":"dog"}`).
		End()
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(body))
	log.Println("body=", body)

	//JSON Post 02
	log.Println("Json Post 02" + HR)
	type BrowserVersionSupport struct {
		Chrome  string
		Firefox string
	}
	ver := BrowserVersionSupport{Chrome: "37.0.2041.6", Firefox: "30.0"}
	request = gorequest.New()
	resp, body, errs = request.Post("http://httpbin.org/post").
		Send(ver).
		End()
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(body))
	log.Println("body=", body)

	//sendMap
	request = gorequest.New()
	resp, body, errs = request.Post("http://httpbin.org/post").
		SendMap(map[string]interface{}{"test01": "123", "test02": 456}).
		End()
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(body))
	log.Println("body=", body)

	//form
	//要加上Type("form")才行，不然是发送的JSON
	log.Println("form" + HR)
	request = gorequest.New()
	resp, body, errs = request.Post("http://httpbin.org/post").
		Type("form").
		Send(ver).
		End()
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(body))
	log.Println("body=", body)

}

//处理错误
func handleErrors(errs []error) {
	if errs != nil {
		for _, v := range errs {
			log.Println("error:", v)
		}
		os.Exit(2)
	}
}
