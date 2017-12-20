package main

import (
	"github.com/mccutchen/go-httpbin/httpbin"
	"net/http"
)

func main() {
	handler := httpbin.NewHTTPBin().Handler()

	http.ListenAndServe(":9999", handler)
}
