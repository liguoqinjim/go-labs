# [colly](https://github.com/gocolly/colly)
安装:`go get github.com/gocolly/colly/v2/`

|实验|简介|说明|
|---|---|---|
|lab001|demo|Collector的使用|
|lab002|demo|修改Collector的配置|
|lab003|example|Debugging|
|lab004|demo||

## NOTICE
### 代替Visit，可以传入更多参数
 ```go
ctx := colly.NewContext()
ctx.Put("filename", filename)
if err := col.Request(http.MethodGet, url, nil, ctx, nil); err != nil {
    log.Fatalf("col.Visit error:%v", err)
}
 ```
### 使用代理
需要使用最新的代码，https://github.com/gocolly/colly/commit/51adaf46046853a4686cad1226431ca972a4a414。  
https://github.com/gocolly/colly/pull/567， 这个代码可以解决使用SetProxyFunc的时候，代理并不是每个请求都会更换的情况