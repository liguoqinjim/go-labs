package routers

import (
	"github.com/kataras/iris/v12"
	"lab150/lab006/controllers"
	"lab150/lab006/middleware"
)

func Register(app *iris.Application) {
	// 路由集使用跨域中间件 CrsAuth()
	// 允许 Options 方法 AllowMethods(iris.MethodOptions)
	main := app.Party("/", middleware.CrsAuth()).AllowMethods(iris.MethodOptions)
	{
		v1 := main.Party("/v1")
		{
			v1.Post("/admin/login", controllers.UserLogin)
			v1.PartyFunc("/admin", func(admin iris.Party) {
				admin.Use(middleware.JwtHandler().Serve, middleware.AuthToken) //登录验证
				admin.Get("/logout", controllers.UserLogout).Name = "退出"
			})
		}
	}
}
