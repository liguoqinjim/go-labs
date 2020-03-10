package controllers

import (
	"github.com/kataras/iris/v12"
	"lab150/lab006/models"
	"net/http"
	"strconv"
)

// 登陆处理程序
func UserLogin(ctx iris.Context) {
	aul := new(models.User)
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(models.Response{Status: false, Msg: nil, Data: "请求参数错误"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	response, status, msg := models.CheckLogin(aul.Username, aul.Password)
	_, _ = ctx.JSON(models.Response{Status: status, Msg: response, Data: msg})
	return
}

// 登出
func UserLogout(ctx iris.Context) {
	aui := ctx.Values().GetString("auth_user_id")
	id, _ := strconv.Atoi(aui)
	models.UserAdminLogout(uint(id))
	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(models.Response{true, nil, "退出"})
}
