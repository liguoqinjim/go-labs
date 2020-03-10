package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	_ "github.com/mattn/go-sqlite3"
	"lab150/lab006/models"
	"lab150/lab006/routers"
)

func NewApp() *iris.Application {
	models.Register() // 数据库初始化
	models.Db.AutoMigrate(
		&models.User{},
		&models.OauthToken{},
	)
	iris.RegisterOnInterrupt(func() {
		_ = models.Db.Close()
	})
	models.CreateUser()

	app := iris.New()
	app.Logger().SetLevel("debug") //设置日志级别
	app.Use(recover.New())
	app.Use(logger.New())
	routers.Register(app) // 注册路由
	return app
}

func main() {
	app := NewApp()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
