package main

import (
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {})
	app.Post("/file", func(ctx iris.Context) {})
	app.Get("/user/{userid:int min(1)}", func(ctx iris.Context) {})

	for n, v := range app.GetRoutes() {
		log.Println(n, v, v.FormattedPath)
	}
}
