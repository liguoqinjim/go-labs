package main

import (
	"github.com/mccutchen/go-httpbin/httpbin"
	"net/http"
)

func main() {
	handler := httpbin.NewHTTPBin().Handler()

	http.ListenAndServe("127.0.0.1:9998", handler)
}
