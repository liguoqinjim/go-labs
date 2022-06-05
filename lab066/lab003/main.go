package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func demo01() {
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "k3",
		Value:  "v3",
		Path:   "/",
		Domain: "httpbin.org",
	}
	cookies = append(cookies, cookie)

	u, _ := url.Parse("http://httpbin.org/cookies")
	jar.SetCookies(u, cookies)
	log.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest("GET", "http://httpbin.org/cookies", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println(string(body))
}

func demo02() {
	//options := cookiejar.Options{
	//	PublicSuffixList: publicsuffix.List,
	//}
	//jar, _ := cookiejar.New(&options)
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "k3",
		Value:  "v3",
		Path:   "/",
		//Domain: "liguoqinjim.cn",
	}
	cookies = append(cookies, cookie)

	u, _ := url.Parse("http://httpbin.liguoqinjim.cn/cookies")
	jar.SetCookies(u, cookies)
	log.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest("GET", "http://httpbin.liguoqinjim.cn/cookies", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Println(string(body))
}

func main() {
	//demo01()
	demo02()
}
