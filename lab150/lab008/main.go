package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.New()

	authConfig := basicauth.Config{
		Users: map[string]string{"admin": "123456"},
	}
	authentication := basicauth.New(authConfig)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/admin")
	})

	needAuth := app.Party("/admin", authentication)
	{
		needAuth.Get("/", h)
		needAuth.Get("/profile", h)
		needAuth.Get("/settings", h)
	}

	return app
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}
