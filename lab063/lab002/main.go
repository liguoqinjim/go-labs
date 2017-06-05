package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/client"
)

func main() {
	var err error

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := cdp.New(ctxt, cdp.WithTargets(client.New().WatchPageTargets(ctxt)))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	//err = c.Run(ctxt, openYunbi())
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Println("打开网站")
	//err = c.Run(ctxt, cdp.Tasks{
	//	cdp.Navigate("https://yunbi.com/?warning=false"),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("打开网站完成")

	//打开sc交易
	err = c.Run(ctxt, openSCMarket())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("输入完成")
}

func openYunbi() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate("https://yunbi.com/signin"),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(`#new_identity > div.ui.stacked.segment > input`, cdp.ByQuery),
		cdp.Click(`#new_identity > div.ui.stacked.segment > input`),
		cdp.WaitVisible(`#wrap > div.ui.text.container.segment > a`, cdp.ByQuery),
		cdp.Click(`#wrap > div.ui.text.container.segment > a`),
	}
}

func openSCMarket() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate("https://yunbi.com/markets/sccny"),
		cdp.WaitVisible(`#new_order_bid > button`, cdp.ByQuery),
		cdp.SendKeys(`#order_bid_origin_volume`, `2222`),
	}
}

func googleSearch(q, text string, site, res *string) cdp.Tasks {
	var buf []byte
	sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
	return cdp.Tasks{
		cdp.Navigate(`https://www.google.com`),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(`#hplogo`, cdp.ByID),
		cdp.SendKeys(`#lst-ib`, q+"\n", cdp.ByID),
		cdp.WaitVisible(`#res`, cdp.ByID),
		cdp.Text(sel, res),
		cdp.Click(sel),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(`#footer`, cdp.ByQuery),
		cdp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, cdp.ByQuery),
		cdp.Location(site),
		cdp.Screenshot(`#testimonials`, &buf, cdp.ByID),
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			return ioutil.WriteFile("testimonials.png", buf, 0644)
		}),
	}
}
