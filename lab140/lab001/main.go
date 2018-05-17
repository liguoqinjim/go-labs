package main

import (
	"github.com/gocolly/colly"
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

	c.OnHTML(`.ip_text`, func(e *colly.HTMLElement) {
		log.Println("OnHTML:value=", e.Text)
	})

	c.OnScraped(func(resp *colly.Response) {
		log.Println("OnScraped:", resp.StatusCode)
	})

	c.Visit("https://www.ipip.net/")
}
