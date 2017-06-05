package main

import (
	"context"
	"log"
	"os"
	"time"

	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/kb"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var val1, val2, val3, val4 string
	err = c.Run(ctxt, sendkeys(&val1, &val2, &val3, &val4))
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

	log.Printf("#input1 value: %s", val1)
	log.Printf("#textarea1 value: %s", val2)
	log.Printf("#input2 value: %s", val3)
	log.Printf("#select1 value: %s", val4)
}

func sendkeys(val1, val2, val3, val4 *string) cdp.Tasks {
	//因为使用了vendor，所以这里要计算一下vendor的路径
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	finalPath := "file:" + path + "/../vendor/github.com/knq/chromedp/testdata/visible.html"

	return cdp.Tasks{
		//cdp.Navigate("file:" + os.Getenv("GOPATH") + "/src/github.com/knq/chromedp/testdata/visible.html"),
		cdp.Navigate(finalPath),
		cdp.WaitVisible(`#input1`, cdp.ByID),
		cdp.WaitVisible(`#textarea1`, cdp.ByID),
		cdp.SendKeys(`#textarea1`, kb.End+"\b\b\n\naoeu\n\ntest1\n\nblah2\n\n\t\t\t\b\bother box!\t\ntest4", cdp.ByID),
		cdp.Value(`#input1`, val1, cdp.ByID),
		cdp.Value(`#textarea1`, val2, cdp.ByID),
		cdp.SetValue(`#input2`, "test3", cdp.ByID),
		cdp.Value(`#input2`, val3, cdp.ByID),
		cdp.SendKeys(`#select1`, kb.ArrowDown+kb.ArrowDown, cdp.ByID),
		cdp.Value(`#select1`, val4, cdp.ByID),
		cdp.Sleep(30 * time.Second),
	}
}
