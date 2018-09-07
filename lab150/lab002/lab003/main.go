package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	//加载模板
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	// GET http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})

	// POST http://localhost:8080/form_action
	app.Post("/form_action", func(ctx iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}

		ctx.Writef("Visitor: %#v", visitor)
	})

	// POST http://localhost:8080/post_value
	app.Post("/post_value", func(ctx iris.Context) {
		username := ctx.PostValueDefault("Username", "iris")
		ctx.Writef("Username:%s", username)
	})

	app.Run(iris.Addr(":8080"))
}

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}
