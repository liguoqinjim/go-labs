package main

import (
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	//redirects to /admin without basic auth
	e.GET("/").Expect().Status(httptest.StatusUnauthorized)
	e.GET("/admin").Expect().Status(httptest.StatusUnauthorized)

	//with valid basic auth
	e.GET("/admin").WithBasicAuth("admin", "123456").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin admin:123456")
	e.GET("/admin/profile").WithBasicAuth("admin", "123456").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/profile admin:123456")
	e.GET("/admin/settings").WithBasicAuth("admin", "123456").Expect().
		Status(httptest.StatusOK).Body().Equal("admin/settings admin:123456")

	//with invalid basic auth
	e.GET("/admin/settings").WithBasicAuth("admin", "admin").Expect().
		Status(httptest.StatusUnauthorized)
}
