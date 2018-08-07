package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	resp, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatalf("http.Get error:%v", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("goquery.NewDocumentFromReader error:%v", err)
	}

	//寻找节点
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		log.Printf("Review %d: %s - %s\n", i, band, title)
	})

	//找一个没有的节点
	n := doc.Find("liguoqinjim")
	if n == nil {
		log.Println("n is nil")
	} else {
		log.Println("n is not nil")
		log.Println("n.Length()=", n.Length())
	}
	//判断这个n的children是否为nil
	if n.Children() == nil {
		log.Println("n.Children() is nil")
	} else {
		log.Println("n.Children() is not nil")
		log.Println("n.Children().Length()=", n.Children().Length())
	}
}
