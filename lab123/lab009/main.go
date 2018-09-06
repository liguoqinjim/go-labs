package main

import (
	"github.com/sclevine/agouti"
	"log"
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
	if err := page.Navigate("http://www.w3school.com.cn/html/html_iframe.asp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	f := page.Find("#intro > iframe")
	if err := f.SwitchToFrame(); err != nil {
		log.Fatalf("f.SwitchToFrame error:%v", err)
	} else {
		log.Printf("swtichtoFrame success")
	}

	if err := page.Find("#course > ul:nth-child(2) > li:nth-child(10) > a").Click(); err != nil {
		log.Printf("Failed to click:%v", err)
	} else {
		log.Printf("Click success")
	}

	//截图
	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab009/tmp/chrome.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	} else {
		log.Printf("screenshot success")
	}
}
