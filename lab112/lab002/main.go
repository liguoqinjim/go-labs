package main

import (
	"log"
	"net/http"
)

func setCookieHandle(w http.ResponseWriter, req *http.Request) {
	c1 := http.Cookie{Name: "key1", Value: "value1"}
	c2 := http.Cookie{Name: "key2", Value: "value2"}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func delCookieHandle(w http.ResponseWriter,req *http.Request){
	
}

func main() {
	http.HandleFunc("/setCookie", setCookieHandle)
	log.Println("Listening...")

	http.ListenAndServe("127.0.0.1:9999", nil)
}
