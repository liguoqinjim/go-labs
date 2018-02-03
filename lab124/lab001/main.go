package main

import (
	"github.com/juju/persistent-cookiejar"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	//查看默认保存位置
	filePath := cookiejar.DefaultCookieFile()
	log.Println("默认保存位置:", filePath)

	//创建一个新的cookieJar，并制定保存的文件
	jar, err := cookiejar.New(&cookiejar.Options{
		Filename: "E:/Workspace/go-labs/src/lab124/lab001/tmp/cookie", //自定义保存文件
	})
	if err != nil {
		log.Fatalf("cookieJar.New error:%v", err)
	}

	//url
	u := &url.URL{Host: "baidu.com", Scheme: "http"}

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
