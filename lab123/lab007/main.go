package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	proxy := agouti.ProxyConfig{
		ProxyType: "manual",
		HTTPProxy: "58.19.81.193:18118",
		SSLProxy:  "58.19.81.193:18118",
	}

	c := agouti.NewCapabilities().Proxy(proxy)
	driver := agouti.PhantomJS(agouti.Desired(c))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer func() {
		//time.Sleep(time.Hour)
		driver.Stop()
	}()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//打开页面
	if err := page.Navigate("http://httpbin.org/ip"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab007/tmp/phantomjs_ipip.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}
}
