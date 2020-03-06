package middleware

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"lab150/lab005/config"
	"net/http"
	"path"
	"time"
)

func LoggerHandler(ctx iris.Context) {
	p := ctx.Request().URL.Path
	method := ctx.Request().Method
	start := time.Now()
	fields := make(map[string]interface{})
	fields["title"] = "访问日志"
	fields["fun_name"] = path.Join(method, p)
	fields["ip"] = ctx.Request().RemoteAddr
	fields["method"] = method
	fields["url"] = ctx.Request().URL.String()
	fields["proto"] = ctx.Request().Proto
	//fields["header"] = ctx.Request().Header
	fields["user_agent"] = ctx.Request().UserAgent()
	fields["x_request_id"] = ctx.GetHeader("X-Request-Id")

	// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = ioutil.NopCloser(buf)
			fields["content_length"] = ctx.GetContentLength()
			fields["body"] = string(body)
		}
	}
	config.Log.Debug("before next")
	ctx.Next()
	config.Log.Debug("after next")

	//下面是返回日志
	fields["res_status"] = ctx.ResponseWriter().StatusCode()
	if ctx.Values().GetString("out_err") != "" {
		fields["out_err"] = ctx.Values().GetString("out_err")
	}
	fields["res_length"] = ctx.ResponseWriter().Header().Get("size")
	if v := ctx.Values().Get("res_body"); v != nil {
		if b, ok := v.([]byte); ok {
			fields["res_body"] = string(b)
		}
	}
	token := ctx.Values().Get("jwt")
	if token != nil {
		fields["uid"] = token.(*jwt.Token).Claims
	}
	timeConsuming := time.Since(start).Nanoseconds() / 1e6
	msg := fmt.Sprintf("[http] %s-%s-%s-%d(%dms)",
		p, ctx.Request().Method, ctx.Request().RemoteAddr, ctx.ResponseWriter().StatusCode(), timeConsuming)
	config.Log.Debug(fields)
	config.Log.Infof(msg)
}