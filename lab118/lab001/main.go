package main

import (
	"github.com/getsentry/raven-go"
	"github.com/henrylee2cn/faygo/errors"
	"os"
)

func init() {
	raven.SetDSN("https://a6699ae2a5ae4cf6abfe505291950044:4c9bc323735b48b2b9466e612afe9089@sentry.io/273021")
}

func main() {
	//错误1
	f, err := os.Open("1.txt")
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
	_ = f

	//错误2
	a := 5
	if a > 3 {
		raven.CaptureErrorAndWait(errors.New("a > 3"), nil)
	}
}
