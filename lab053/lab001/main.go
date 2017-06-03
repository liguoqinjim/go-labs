package main

import (
	"github.com/hoisie/web"
)

func hello(val string) string {
	return "hello " + val
}

func hi(ctx *web.Context, val string) {
	ctx.WriteString("hi " + val)
}

func main() {
	web.Get("/hello(.*)", hello)
	web.Get("/hi(.*)", hi)

	web.Run("0.0.0.0:9999")
}
