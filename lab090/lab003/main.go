package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	HR = "------------------------------------------------------------------"
)

func main() {
	//EndBytes
	log.Println("EndBytes" + HR)
	resp, bodyBytes, errs := gorequest.New().Get("http://httpbin.org/headers").EndBytes()
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("body.length=", len(bodyBytes))

	//EndStruct
	log.Println("EndStruct" + HR)
	type responseStruct struct {
		Headers struct {
			Accept                  string `json:"Accept"`
			AcceptEncoding          string `json:"Accept-Encoding"`
			AcceptLanguage          string `json:"Accept-Language"`
			Connection              string `json:"Connection"`
			Cookie                  string `json:"Cookie"`
			Host                    string `json:"Host"`
			UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
			UserAgent               string `json:"User-Agent"`
		} `json:"headers"`
	}
	var rs responseStruct
	resp, _, errs = gorequest.New().Get("http://httpbin.org/headers").EndStruct(&rs)
	handleErrors(errs)
	log.Println("resp.Status=", resp.Status)
	log.Println("rs=", rs)

	//Timeout
	log.Println("Timeout" + HR)
	request := gorequest.New().Timeout(2 * time.Millisecond)
	resp, body, errs := request.Get("http://google.com").End()
	if errs != nil {
		log.Println(errs)
	}
	log.Println("resp=", resp)
	log.Println("body=", body)

	//Retry
	log.Println("Retry" + HR)
	request = gorequest.New()
	resp, body, errs = request.Get("http://google.com/").
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		End()
	if errs != nil {
		log.Println(errs)
	}
	log.Println("resp=", resp)
	log.Println("body=", body)
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
