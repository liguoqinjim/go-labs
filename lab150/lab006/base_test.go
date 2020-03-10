package main

import (
	"flag"
	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"os"
	"testing"
)

const baseUrl = "/v1/admin/"       // 接口地址
const loginUrl = baseUrl + "login" // 登陆接口地址
var (
	app *iris.Application
)

//单元测试基境
func TestMain(m *testing.M) {
	// 初始化app
	app = NewApp()
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

// 单元测试 login 方法
func login(t *testing.T, Object interface{}, StatusCode int, Status bool, Msg string) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: true})
	e.POST(loginUrl).WithJSON(Object).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg)
	return
}
