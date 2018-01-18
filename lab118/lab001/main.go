package main

import (
	"errors"
	"github.com/getsentry/raven-go"
	"os"
)

func init() {
	raven.SetDSN("https://a6699ae2a5ae4cf6abfe505291950044:4c9bc323735b48b2b9466e612afe9089@sentry.io/273021")
}

func main() {
	//用户信息
	raven.SetUserContext(&raven.User{ID: "123", Email: "136542728@qq.com"})

	//错误1
	f, err := os.Open("1.txt")
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	_ = f

	//错误2(添加标签)
	a := 10
	if a > 6 {
		raven.CaptureErrorAndWait(errors.New("a > 6"), map[string]string{"myErrorCode": "999001"})
	}
}
