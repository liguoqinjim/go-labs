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
	if err := page.Navigate("http://www.runoob.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//点击Click
	if err := page.Find("body > div.container.main > div > div.col.middle-column-home > div.codelist.codelist-desktop.cate1 > a:nth-child(2)").Click(); err != nil {
		log.Printf("Failed to click:%v", err)
	}

	if err := page.Find("#leftcolumn > a:nth-child(12)").Click(); err != nil {
		log.Printf("Failed to click:%v", err)
	}

	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab008/tmp/chrome.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}
}
