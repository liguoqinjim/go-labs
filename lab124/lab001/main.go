package main

import (
	"github.com/juju/persistent-cookiejar"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//查看默认保存位置
	filePath := cookiejar.DefaultCookieFile()
	log.Println("filePath=", filePath)

	//创建一个新的cookieJar
	jar, err := cookiejar.New(&cookiejar.Options{
		Filename: "E:/Workspace/go-labs/src/lab124/lab001/tmp/cookie",
	})
	if err != nil {
		log.Fatalf("cookieJar.New error:%v", err)
	}
	//set cookie
	u := &url.URL{Host: "baidu.com", Scheme: "http"}
	c1 := &http.Cookie{
		Name:     "first_cookie",
		Value:    "vanyar",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour),
	}
	cookies := []*http.Cookie{c1}
	jar.SetCookies(u, cookies)

	//get cookie
	cs := jar.AllCookies()
	for n, v := range cs {
		log.Printf("getCookie[%d]:%+v", n, v)
	}

	//保存cookie
	err = jar.Save()
	if err != nil {
		log.Printf("jar.Save error:%v", err)
	} else {
		log.Println("save success")
	}
}
