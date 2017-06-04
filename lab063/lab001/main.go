package main

import (
	"context"
	"log"
	"time"

	cdp "github.com/knq/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt, cdp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	s := "test1"
	err = c.Run(ctxt, click(&s))
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

func click(res *string) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(`https://www.baidu.com/?tn=95932978_hao_pg`),
		//cdp.WaitVisible(`#footer`),
		cdp.Text(`#wd`, res, cdp.NodeVisible, cdp.ByID),
		cdp.Click(`#su`, cdp.NodeVisible),
		cdp.Sleep(150 * time.Second),
	}
}
