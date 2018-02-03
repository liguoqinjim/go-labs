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

	//打开页面
	if err := page.Navigate("https://www.baidu.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	//得到cookie
	cookies, err := page.GetCookies()
	if err != nil {
		log.Printf("page.GetCookies error:%v", err)
	}
	for n, v := range cookies {
		log.Printf("Cookie[%d],%+v", n, v)
	}

	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab004/tmp/chrome_baidu.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}

	//设置cookie
	for i := 1; i <= 10; i++ {
		c := &http.Cookie{
			Name:     fmt.Sprintf("key%d", i),
			Value:    fmt.Sprintf("value%d", i),
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
		}
		if err := page.SetCookie(c); err != nil {
			log.Printf("page.SetCookie error:%v", err)
		}

		time.Sleep(time.Second * 10)
	}
}
