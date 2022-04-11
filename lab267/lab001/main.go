package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			log.Println("add header")
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})

	if err := http.ListenAndServe(":38080", proxy); err != nil {
		log.Fatalf("err:%v", err)
	}
}
