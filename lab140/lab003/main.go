package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"log"
)

func main() {
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	if err := c.Visit("http://httpbin.org/user-agent"); err != nil {
		log.Fatalf("c.Visit error:%v", err)
	}
}
