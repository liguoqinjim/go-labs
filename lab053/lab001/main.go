package main

import (
	"github.com/hoisie/web"
)

func hello(val string) string { //这个函数的参数数量是要和router中的参数数量是一样的
	return "hello " + val
}

func hi(ctx *web.Context, val string) {
	ctx.WriteString("hi " + val)
}

func testHead(ctx *web.Context, val string) string {
	ctx.SetHeader("X-Powered-By", "web.go", true)
	ctx.SetHeader("X-Frame-Options", "DENY", true)
	ctx.SetHeader("Connection", "close", true)
	return "testhead " + val
}

func main() {
	web.Get("/hello(.*)", hello)
	web.Get("/hi(.*)", hi)
	web.Get("/testhead(.*)", testHead)

	web.Run("0.0.0.0:9999")
}
