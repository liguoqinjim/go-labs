package main

import (
	"context"
	"log"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/client"
	"os"
	"time"
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
	err = c.Run(ctxt, baiduSearch())
	if err != nil {
		log.Fatal(err)
	}
}

func baiduSearch() cdp.Tasks {
	//file:///E:/Workspace/go-labs/src/lab063/lab003/lab001/testdata/time.html
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := "file:///" + pwd + "/testdata/time.html"

	return cdp.Tasks{
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			log.Println("time1", time.Now().Unix())
			log.Println("time1", time.Now().UnixNano())
			log.Println(filePath)
			return nil
		}),
		cdp.Navigate(filePath),
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			log.Println("time2", time.Now().Unix())
			log.Println("time2", time.Now().UnixNano())
			return nil
		}),
	}
}
