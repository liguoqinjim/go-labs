package main

import (
	"github.com/getsentry/raven-go"
)

func init() {
	raven.SetDSN("https://a6699ae2a5ae4cf6abfe505291950044:4c9bc323735b48b2b9466e612afe9089@sentry.io/273021")
}

func main() {
	raven.CapturePanicAndWait(func() {
		panic(99999)
	}, nil)
}
