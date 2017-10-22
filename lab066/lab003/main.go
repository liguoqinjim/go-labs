package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
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
