package main

import (
	"github.com/henrylee2cn/surfer"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	HR = "------------------------------------------------------------------"
)

func main() {
	//默认surf内核
	log.Println("默认surf内核" + HR)
	resp, err := surfer.Download(&surfer.Request{
		Url:          "http://httpbin.org/user-agent",
		Method:       http.MethodGet,
		DownloaderID: 0,
	})
	handleError(err)
	log.Println("resp.Status")

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))

	//phamtomjs内核
	log.Println("phamtomjs内核" + HR)
	resp, err = surfer.Download(&surfer.Request{
		Url:          "http://httpbin.org/user-agent",
		Method:       http.MethodGet,
		DownloaderID: 1,
	})
	handleError(err)

	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))

	surfer.DestroyJsFiles()

	time.Sleep(10e9)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
