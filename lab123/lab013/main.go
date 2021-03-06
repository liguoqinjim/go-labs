package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//打开url
	if err := page.Navigate("https://www.baidu.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//截图
	if err := page.Screenshot("tmp/chrome_baidu.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}
}
