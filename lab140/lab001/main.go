package main

import (
	"github.com/gocolly/colly/v2"
	"log"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Println("OnRequest:", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("OnError:", r.StatusCode, e)
	})

	c.OnResponse(func(resp *colly.Response) {
		log.Println("OnResponse:", resp.StatusCode)
	})

	c.OnHTML(`body > p:nth-child(1)`, func(e *colly.HTMLElement) {
		log.Println("OnHTML:value=", e.Text)
	})

	c.OnScraped(func(resp *colly.Response) {
		log.Println("OnScraped:", resp.StatusCode)
	})

	if err := c.Visit("https://2020.ip138.com/"); err != nil {
		log.Fatalf("c.Visit error:%v", err)
	}
}
