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
	if err := page.Navigate("https://www.baidu.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	if err := page.Navigate("https://www.baidu.com/s?wd=1"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//time.Sleep(time.Second * 10)

	cnt, _ := page.WindowCount()
	log.Println("window cnt=", cnt)

	if err := page.NextWindow(); err != nil {
		log.Fatalf("NextWindow error:%v", err)
	}

	page2, err := agouti.JoinPage()
	if err != nil {
		log.Fatalf("NewPage error:%v", err)
	}
	_ = page2

}
