package main

import (
	"github.com/kataras/iris/v12"
	"testing"
)

//登陆成功
func TestUserLoginSuccess(t *testing.T) {
	oj := map[string]string{
		"username": "username",
		"password": "password",
	}
	login(t, oj, iris.StatusOK, true, "登陆成功")
}

// 输入不存在的用户名登陆
func TestUserLoginWithErrorName(t *testing.T) {
	oj := map[string]string{
		"username": "err_user",
		"password": "password",
	}
	login(t, oj, iris.StatusOK, false, "用户不存在")
}

// 输入错误的登陆密码
func TestUserLoginWithErrorPwd(t *testing.T) {
	oj := map[string]string{
		"username": "username",
		"password": "admin",
	}
	login(t, oj, iris.StatusOK, false, "用户名或密码错误")
}

// 输入登陆密码格式错误
func TestUserLoginWithErrorFormtPwd(t *testing.T) {
	oj := map[string]string{
		"username": "username",
		"password": "123",
	}
	login(t, oj, iris.StatusOK, false, "密码格式错误")
}

// 输入登陆密码格式错误
func TestUserLoginWithErrorFormtUserName(t *testing.T) {
	oj := map[string]string{
		"username": "df",
		"password": "123",
	}
	login(t, oj, iris.StatusOK, false, "用户名格式错误")
}
