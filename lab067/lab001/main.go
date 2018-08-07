package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ExampleScrape() {
	resp, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatalf("http.Get error:%v", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("goquery.NewDocumentFromReader error:%v", err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		log.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func main() {
	ExampleScrape()
}
