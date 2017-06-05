package main

import (
	"context"
	"log"
	"time"

	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/runner"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	//最后的参数是设置代理，因为给出的例子里面都是需要用代理才能访问的
	c, err := cdp.New(ctxt, cdp.WithLog(log.Printf), cdp.WithRunnerOptions(runner.Proxy("127.0.0.1:1080")))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	err = c.Run(ctxt, click())
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func click() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(`https://golang.org/pkg/time/`),
		cdp.WaitVisible(`#footer`),
		//模拟点击
		cdp.Click(`#pkg-overview`, cdp.NodeVisible),
		cdp.Sleep(150 * time.Second),
	}
}
