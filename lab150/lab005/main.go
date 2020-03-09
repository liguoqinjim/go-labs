package main

import (
	"github.com/kataras/iris/v12"
	"lab150/lab005/config"
	"lab150/lab005/middleware"
)

func main() {
	app := iris.New()

	//logger_middleware
	app.Use(middleware.RequestId)
	//app.Use(middleware.LoggerHandler)
	app.Use(middleware.LoggerHandler2)


	app.Handle("GET", "/", func(ctx iris.Context) {
		config.Log.Debug("handling...")
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
