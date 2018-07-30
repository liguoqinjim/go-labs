package main

import (
	"io/ioutil"
	"github.com/henrylee2cn/surfer"
	"log"
)

//代理
const (
	HTTP_PROXY = "http://114.215.174.227:8080"
	SS_PROXY   = "http://127.0.0.1:1080"
)

func main() {
	resp, err := surfer.Download(&surfer.Request{
		Url: "http://ip.cip.cc/",
	})
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("resp.Body: %s\n", b)

	//使用http_proxy
	if HTTP_PROXY != "" {
		resp, err = surfer.Download(&surfer.Request{
			Url:   "http://ip.cip.cc/",
			Proxy: HTTP_PROXY,
		})

		if err != nil {
			log.Fatal(err)
		}

		b, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		log.Printf("resp.Body: %s\n", b)
	}

	//使用ss代理
	if SS_PROXY != "" {
		resp, err = surfer.Download(&surfer.Request{
			Url:   "http://ip.cip.cc/",
			Proxy: SS_PROXY,
		})

		if err != nil {
			log.Fatal(err)
		}

		b, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		log.Printf("resp.Body: %s\n", b)
	}
}
