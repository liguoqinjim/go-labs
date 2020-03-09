package core

import "github.com/kataras/iris/v12"

func GetReqID(ctx iris.Context) string {
	requestId := ctx.Values().GetString("req-id")

	return requestId
}
