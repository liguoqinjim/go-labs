package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
)

const (
	HTTP_PROXY = "http://114.215.174.227:8080"
	SS_PROXY   = "http://127.0.0.1:1080"
)

func main() {
	request := gorequest.New()
	_, body, errs := request.Get("http://ip.cip.cc").End()
	handleErrors(errs)
	log.Println("body=", body)

	//http_proxy
	if HTTP_PROXY != "" {
		_, body, errs = request.Get("http://ip.cip.cc").Proxy(HTTP_PROXY).End()
		handleErrors(errs)
		log.Println("body=", body)
	}

	//ss_proxy
	if SS_PROXY != "" {
		_, body, errs = request.Get("http://ip.cip.cc").Proxy(SS_PROXY).End()
		handleErrors(errs)
		log.Println("body=", body)
	}
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
