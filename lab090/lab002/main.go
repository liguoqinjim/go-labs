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
	//Callback
	log.Println("Callback" + HR)
	request := gorequest.New()
	request.Get("http://httpbin.org/headers").End(printStatus)

	//Multipart/Form-Data
	log.Println("Multipart/Form-Data" + HR)
	gorequest.New().Post("http://httpbin.org/post").
		Type("multipart").
		Send(`{"query1":"test"}`).
		End(printStatus)

	//Proxy
	log.Println("Proxy" + HR)
	gorequest.New().Proxy("http://103.8.194.6:53281").
		Get("http://httpbin.org/get").End(printStatus)

	//Basic Authentication
	log.Println("Basic Authentication" + HR)
	gorequest.New().
		SetBasicAuth("username", "password").
		Get("http://httpbin.org/get").
		End(printStatus)
}

func printStatus(resp gorequest.Response, body string, errs []error) {
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
