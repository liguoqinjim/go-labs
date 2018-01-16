package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	urlProxy, _ := url.Parse("https://94.177.250.95:80")
	c := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlProxy),
		},
	}
	if resp, err := c.Get("http://httpbin.org/ip"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
