package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gorequest.New()

	r.Timeout(time.Second * 2)
	//statusCode一定要指定，没有默认值
	r.Retry(3, time.Second*2, http.StatusInternalServerError)

	resp, body, errs := r.Get("http://httpbin.org/status/500").EndBytes()
	if errs != nil {
		log.Fatalf("get error:%v", errs)
	}
	defer resp.Body.Close()

	log.Printf("body:%s", body)
	log.Println(resp.StatusCode)
}
