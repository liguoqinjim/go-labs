package main

import (
	"github.com/henrylee2cn/surfer"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	HR = "------------------------------------------------------------------"
)

func main() {
	//Get
	log.Println("Get" + HR)
	resp, err := surfer.Download(&surfer.Request{
		Url:    "http://httpbin.org/get?a=1",
		Method: http.MethodGet,
	})
	handleError(err)
	log.Println("resp.Status=", resp.Status)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))

	//Post
	log.Println("Post" + HR)
	values1, err := url.ParseQuery("s1=v1&s2=v2")
	handleError(err)
	form1 := surfer.Form{
		Values: values1,
	}
	resp, err = surfer.Download(&surfer.Request{
		Url:    "http://httpbin.org/post",
		Method: http.MethodPost,
		Body:   form1,
	})
	handleError(err)
	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println("body=", string(b))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
