package main

import (
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func main() {
	proxy := agouti.ProxyConfig{
		ProxyType: "manual",
		HTTPProxy: "http://127.0.0.1:1080",
		SSLProxy:  "http://127.0.0.1:1080",
	}

	c := agouti.NewCapabilities().Proxy(proxy)
	driver := agouti.ChromeDriver(agouti.Desired(c))
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
	if err := page.Navigate("http://www.cip.cc/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
}
