# [colly](https://github.com/gocolly/colly)
安装:`go get github.com/gocolly/colly/v2/`

|实验|简介|说明|
|---|---|---|
|lab001|demo|Collector的使用|
|lab002|demo|修改Collector的配置|
|lab003|example|Debugging|
|lab004|demo||

## NOTICE
 - 代理Visit，可以传入更多参数
 ```go
ctx := colly.NewContext()
ctx.Put("filename", filename)
if err := col.Request(http.MethodGet, url, nil, ctx, nil); err != nil {
    log.Fatalf("col.Visit error:%v", err)
}
 ```