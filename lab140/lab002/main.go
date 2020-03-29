package main

import (
	"github.com/gocolly/colly/v2"
	"log"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36"
	url       = "http://httpbin.org/user-agent"
)

func main() {
	//colly有默认配置的user-agent
	c1 := colly.NewCollector()
	c1.OnResponse(func(resp *colly.Response) {
		log.Printf("c1.resp:\n%s", resp.Body)
	})
	if err := c1.Visit(url); err != nil {
		log.Fatalf("c1.Visit error:%v", err)
	}

	//自定义设置user agent
	c2 := colly.NewCollector(
		colly.UserAgent(userAgent),
		colly.AllowURLRevisit(),
	)
	c2.OnResponse(func(resp *colly.Response) {
		log.Printf("c2.resp:\n%s", resp.Body)
	})
	if err := c2.Visit(url); err != nil {
		log.Fatalf("c2.Visit error:%v", err)
	}

	//任何时候都可以修改collector的配置
	//NOTICE:这一次Visit调用是无效的，因为我们visit是和第一次同样的url，
	//AllowURLRevisit默认为false，false的时候colly默认是不会访问相同的url的，还会报错
	if err := c1.Visit(url); err != nil {
		if err == colly.ErrAlreadyVisited {
			//URL already visited error
			log.Println("URL already visited error")
		} else {
			log.Fatalf("c1.Visit error:%v", err)
		}
	}

	//对c1的配置进行修改
	c1.AllowURLRevisit = true
	c1.UserAgent = userAgent
	if err := c1.Visit(url); err != nil {
		log.Fatalf("c1.Visit error:%v", err)
	}
}
