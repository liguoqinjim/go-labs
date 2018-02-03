package main

import (
	"fmt"
	"github.com/juju/persistent-cookiejar"
	"github.com/sclevine/agouti"
	"log"
	"net/url"
	"strconv"
	"time"
)

func main() {
	//url
	u := &url.URL{Host: "httpbin.org", Scheme: "http"}

	//创建cookieJar
	jar, err := cookiejar.New(&cookiejar.Options{
		Filename: "E:/Workspace/go-labs/src/lab124/lab002/tmp/cookies",
	})
	if err != nil {
		log.Fatalf("cookiejar.New error:%v", err)
	}

	//判断保存的cookie
	cookiesOrigin := jar.AllCookies()
	log.Println("cookiesOrigin length=", len(cookiesOrigin))

	//判断当前最大的key
	keyMax := 0
	for n, v := range cookiesOrigin {
		log.Printf("cookiesOrigin[%d]:%+v", n, v)

		key, _ := strconv.Atoi(v.Name)
		if key > keyMax {
			keyMax = key
		}
	}
	keyMax++

	//访问httpbin.org
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//先访问example.com
	if err := page.Navigate("http://example.com"); err != nil {
		log.Fatalf("page.Navigate error:%v", err)
	}

	//设置page的cookies
	for _, v := range cookiesOrigin {
		page.SetCookie(v)
	}

	//查看page的cookies是否设置成功
	cookiesPage, err := page.GetCookies()
	if err != nil {
		log.Printf("page.GetCookies error:%v", err)
	} else {
		log.Println("page.GetCookies success:length=", len(cookiesPage))
	}

	for n, v := range cookiesPage {
		log.Printf("cookiesPage[%d]:%+v", n, v)
	}

	if err := page.Navigate(fmt.Sprintf("http://httpbin.org/cookies/set?%d=%d", keyMax, keyMax)); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	if err := page.Screenshot("E:/Workspace/go-labs/src/lab124/lab002/tmp/chrome_httpbin.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}

	//得到cookie
	cookies, err := page.GetCookies()
	if err != nil {
		log.Printf("page.GetCookies error:%v", err)
	}
	for n, v := range cookies {
		log.Printf("Cookie[%d],%+v", n, v)
		//重置一下expire时间，不然这个库会过滤
		v.Expires = time.Now().Add(time.Hour)
	}
	jar.SetCookies(u, cookies)

	//jar保存
	jar.Save()
}
