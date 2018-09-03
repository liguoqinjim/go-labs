package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"net/http"
	"time"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer func() {
		time.Sleep(time.Hour)
		driver.Stop()
	}()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//先打开example.com
	if err := page.Navigate("http://example.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//打开example之后再，设置cookie
	for i := 1; i <= 3; i++ {
		c := &http.Cookie{
			Name:     fmt.Sprintf("key%d", i),
			Value:    fmt.Sprintf("value%d", i),
			Domain:   "www.baidu.com",
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
		}
		if err := page.SetCookie(c); err != nil {
			log.Printf("page.SetCookie error:%v", err)
		} else {
			log.Println("page.SetCookie success")
		}
	}

	if err := page.Navigate("http://www.baidu.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//读取cookies
	cookiesPage, err := page.GetCookies()
	if err != nil {
		log.Printf("page.GetCookies error:%v", err)
	}
	for n, v := range cookiesPage {
		log.Printf("cookiesPage[%d],%+v", n, v)
	}
}
