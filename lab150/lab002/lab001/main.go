package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
