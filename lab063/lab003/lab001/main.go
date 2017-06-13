package main

import (
	"context"
	"log"

	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/client"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome
	c, err := cdp.New(ctxt, cdp.WithTargets(client.New().WatchPageTargets(ctxt)))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var site, res string
	err = c.Run(ctxt, baiduSearch())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("saved screenshot of #testimonials from search result listing `%s` (%s)", res, site)
}

func baiduSearch() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(`https://www.baidu.com`),
	}
}
