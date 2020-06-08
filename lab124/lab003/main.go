package main

import (
	cookiejar "github.com/juju/persistent-cookiejar"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	//url
	u := &url.URL{Host: "bilibili.com", Scheme: "http"}

	jar, err := cookiejar.New(&cookiejar.Options{
		Filename: "./tmp/cookies",
	})
	if err != nil {
		log.Fatalf("cookiejar.New error:%v", err)
	}

	cookies := jar.Cookies(u)
	c := "_uuid=123; udjc=456; cdnajk=16; sid=cda12cda; w21=w21za; cdas=cdsa"
	cs := strings.Split(c, ";")
	for _, v := range cs {
		vs := strings.Split(v, "=")
		log.Println(vs)
		cookies = append(cookies, &http.Cookie{
			Name:     vs[0],
			Value:    vs[1],
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 24 * 365),
		})
	}
	jar.SetCookies(u, cookies)

	if err := jar.Save(); err != nil {
		log.Fatalf("jar save error:%v", err)
	}
}
