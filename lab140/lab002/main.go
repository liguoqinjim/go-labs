package main

import (
	"github.com/gocolly/colly"
	"log"
)

const chrome_user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36"
const visit_url = "http://httpbin.org/user-agent"

func main() {
	//默认配置的user-agent
	c1 := colly.NewCollector()
	c1.OnResponse(func(resp *colly.Response) {
		log.Printf("c1.resp:\n%s", resp.Body)
	})

	c1.Visit(visit_url)

	//自定义配置
	c2 := colly.NewCollector(
		colly.UserAgent(chrome_user_agent),
		colly.AllowURLRevisit(),
	)
	c2.OnResponse(func(resp *colly.Response) {
		log.Printf("c2.resp:\n%s", resp.Body)
	})

	c2.Visit(visit_url)

	//任何时候都可以修改collector的配置
	//注意：这一次Visit调用是无效的，因为我们visit是和第一次同样的url
	//colly默认是不会访问一样的url的
	c1.Visit(visit_url)

	//对c1的配置进行修改
	c1.AllowURLRevisit = true
	c1.UserAgent = chrome_user_agent
	c1.Visit(visit_url)
}
