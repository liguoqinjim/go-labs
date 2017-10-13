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
	request := gorequest.New()

	//查看cookie
	log.Println("查看cookie" + HR)
	request.Get("http://httpbin.org/cookies").End(printStatus)

	//http://httpbin.org/cookies
	log.Println("Get" + HR)
	resp, body, errs := request.Get("http://httpbin.org/cookies/set?k1=v1&k2=v2").End()
	if errs != nil {
		log.Println(errs)
		os.Exit(2)
	}
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(body))
	log.Println("body=", body)

	//查看cookie
	log.Println("查看cookie" + HR)
	//log.Println("Cookie:", request.Cookies)
	request.Get("http://httpbin.org/cookies").End(printStatus)

	//新开一个request
	request2 := gorequest.New()
	//request2.SetDebug(true)

	//查看Cookie
	log.Println("查看Cookie" + HR)
	request2.Get("http://httpbin.org/cookies").End(printStatus)

	log.Println("Get" + HR)
	request2.Get("http://httpbin.org/cookies/set?k1=v1&k2=v2").
		RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
			for attr, val := range via[0].Header {
				if _, ok := req.Header[attr]; !ok {
					req.Header[attr] = val
				}
			}
			log.Println("req.Response.Cookies=", req.Response.Cookies())
			log.Println("request2.Cookies1=", request2.Cookies)
			request2.AddCookies(req.Response.Cookies())
			log.Println("request2.Cookies2=", request2.Cookies)
			return nil
		}).
		End(printStatus)
	log.Println("request2.Cookies3=", request2.Cookies)

	log.Println("查看Cookie" + HR)
	request2.Get("http://httpbin.org/cookies").End(printStatus)
	log.Println("request2.Cookies=", request2.Cookies)
}

func printStatus(resp gorequest.Response, body string, errs []error) {
	handleErrors(errs)
	//log.Println("resp.Status=", resp.Status)
	//log.Println("body.Header=", resp.Header)
	//log.Println("body.length=", len(body))
	//log.Println("body=", body)
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
