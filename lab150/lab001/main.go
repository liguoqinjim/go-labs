package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	//得到路径里的参数 (不会处理/user和/user/，这两个连接的)
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	//不会处理(/user/tom/)
	//POST http://localhost:8080/user/tom/swim
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	//得到参数(/welcome?firstname=Jane&lastname=Doe)
	//http://127.0.0.1:8080/welcome?firstname=jim&lastname=li
	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "defaultName")
		// shortcut for ctx.Request().URL.Query().Get("lastname").
		lastname := ctx.URLParam("lastname")

		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	//form表单
	//POST http://localhost:8080/form_post
	app.Post("/form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//得到referer
	//GET http://localhost:8080/referer
	app.Get("/referer", func(ctx context.Context) {
		r := ctx.GetReferrer()

		switch r.Type {
		case context.ReferrerSearch:
			ctx.Writef("Search %s: %s\n", r.Label, r.Query)
			ctx.Writef("Google: %s\n", r.GoogleType)
		case context.ReferrerSocial:
			ctx.Writef("Social %s\n", r.Label)
		case context.ReferrerIndirect:
			ctx.Writef("Indirect: %s\n", r.URL)
		}
	})

	app.Run(iris.Addr(":8080"))
}
