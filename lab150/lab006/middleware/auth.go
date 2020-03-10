package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
	"lab150/lab006/models"
	"net/http"
	"time"
)

func AuthToken(ctx context.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	token := models.GetOauthTokenByToken(value.Raw) //获取 access_token 信息
	if token.Revoked || token.ExpressIn < time.Now().Unix() {
		//_, _ = ctx.Writef("token 失效，请重新登录") // 输出到前端
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.StopExecution()
		return
	} else {
		ctx.Values().Set("auth_user_id", token.UserId)
	}
	ctx.Next()
}
