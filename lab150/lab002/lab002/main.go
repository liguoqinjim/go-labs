package main

import (
	"github.com/kataras/iris"
)

//用来绑定的struct
type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()

	//读取模板
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))

	//处理错误的handler
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// Values()用来在handler和middleware之间传递
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
	})

	//
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	//POST: scheme://mysubdomain.$domain.com/decode
	app.Subdomain("mysubdomain.").Post("/decode", func(ctx iris.Context) {})

	//POST http://localhost:8080/decode
	app.Post("/decode", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	//GET http://localhost:8080/encode
	app.Get("/encode", func(ctx iris.Context) {
		doe := User{
			Username:  "Johndoe",
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}

		ctx.JSON(doe)
	})

	//GET http://localhost:8080/profile/anytypeofstring
	app.Get("/profile/{username:string}", profileUsername)

	userRoutes := app.Party("/users", logThisMiddleware)
	{
		userRoutes.Get("/{id:int min(1)}", getUserById)
		userRoutes.Post("/create", createUser)
	}

	//listening
	//POST http://localhost:8080/decode
	//GET http://localhost:8080/encode
	//GET http://localhost:8080/profile/anytypeofstring
	//GET http://localhost:8080/users/1
	//POST http://localhost:8080/users/create
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

	// .Next is required to move forward to the chain of handlers,
	// if missing then it stops the execution at this handler.
	ctx.Next()
}

func profileUsername(ctx iris.Context) {
	//Params方法只能得到路径里面的参数
	username := ctx.Params().Get("username")
	//给view值，第一个值Username就是模板里面的值
	ctx.ViewData("Username", username)
	ctx.View("user/profile.html")
}

func getUserById(ctx iris.Context) {
	userID := ctx.Params().Get("id")

	user := User{Username: "username" + userID} // your own db fetch here instead of user :=...

	ctx.XML(user)
}

func createUser(ctx iris.Context) {
	//var user User
	user := User{}
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "creating user, read and parse form failed. "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	//和html/template一样，第一个值是""，那么我们在模板里面使用{{.}}，就表示了这个object
	ctx.ViewData("", user)
	ctx.View("user/create_verification.html")
}
