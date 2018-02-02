package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.Selenium()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("https://www.baidu.com"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab002/tmp/firefox_baidu.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}
}
