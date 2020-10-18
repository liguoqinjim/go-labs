package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"strings"
)

func main() {
	getSetCookie()
}

func getSetCookie() {
	targetUrl := "http://httpbin.org/cookies/set?a=1&b=2&c=3"
	r := gorequest.New().Get(targetUrl)

	resp, _, errs := r.RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
		//转发
		setCookie := req.Response.Header.Get("Set-Cookie")
		log.Println(setCookie)
		setCookies := req.Response.Header.Values("Set-Cookie")
		log.Println(setCookies)
		log.Println(len(setCookies))

		rawCookies := ""
		for n, v := range setCookies {
			vs := strings.Split(v, ";")
			log.Println(vs[0])
			rawCookies += vs[0]
			if n != len(setCookies)-1 {
				rawCookies += ";"
			}
		}

		//解析
		header := http.Header{}
		header.Add("Cookie", rawCookies)
		request := http.Request{Header: header}

		log.Println(request.Cookies()) // [cookie1=value1 cookie2=value2]
		cookies := request.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie.Name, cookie.Value)
		}

		return nil
	}).
		EndBytes()
	if errs != nil {
		log.Fatalf("r.EndBytes error:%v", errs)
	}
	defer resp.Body.Close()
}
