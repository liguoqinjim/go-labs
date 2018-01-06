package main

import (
	"log"
	"net/http"
)

func indexHandle(w http.ResponseWriter, req *http.Request) {
	log.Println("req.Proto=", req.Proto)
	log.Println("req.TLS=", req.TLS)
	log.Println("req.Host=", req.Host)
	log.Println("req.RequestURI=", req.RequestURI)

	scheme := "http://"
	if req.TLS != nil {
		scheme = "https://"
	}

	fullUrl := scheme + req.Host + req.RequestURI
	log.Println("fullUrl=", fullUrl)
}

func main() {
	http.HandleFunc("/index", indexHandle)

	log.Println("Listening...")
	http.ListenAndServe("127.0.0.1:9090", nil)
}
