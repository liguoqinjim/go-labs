package main

import (
	"github.com/juju/persistent-cookiejar"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	u = &url.URL{Host: "baidu.com", Scheme: "http"}
)

func main() {
	//default file
	filePath := cookiejar.DefaultCookieFile()
	log.Println("default filePath:", filePath)

	//创建一个新的cookieJar，并制定保存的文件
	jar, err := cookiejar.New(&cookiejar.Options{
		Filename: "tmp/cookie.json", //自定义保存文件
	})
	if err != nil {
		log.Fatalf("cookieJar.New error:%v", err)
	}

	//查看文件里面是否有cookie
	cookiesOrigin := jar.Cookies(u)
	cookieValueMax := 0
	for n, v := range cookiesOrigin {
		log.Printf("cookieOrigin[%d]:%+v", n, v)
		value, err := strconv.Atoi(v.Value)
		if err != nil {
			log.Printf("cookie value error:%v", err)
		}
		if value > cookieValueMax {
			cookieValueMax = value
		}
	}

	//set cookie
	c1 := &http.Cookie{
		Name:     strconv.Itoa(len(cookiesOrigin) + 1),
		Value:    strconv.Itoa(cookieValueMax + 5),
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 365),
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
