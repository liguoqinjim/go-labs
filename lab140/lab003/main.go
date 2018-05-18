package main

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	c.Visit("http://httpbin.org/user-agent")
}
