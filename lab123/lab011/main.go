package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	//给phantomjs设置user-agent
	c := agouti.NewCapabilities()
	c["phantomjs.page.customHeaders.User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36"

	driver := agouti.PhantomJS(agouti.Desired(c))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs.exe"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//打开url
	if err := page.Navigate("http://httpbin.org/get"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//截图
	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab011/tmp/phantomjs_baidu.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}
}
