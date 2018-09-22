package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//给response添加cookie
func setCookieHandle(w http.ResponseWriter, req *http.Request) {
	//设置到期时间
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "vanyar",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour),
	}

	//不设置到期时间，session型cookie
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "noldor",
		HttpOnly: true,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

//得到request的cookie
func getCookieHandle(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No flash cookie found")
		}
	} else {
		fmt.Fprintln(w, "flash=", c)
	}

	cs := req.Cookies()
	fmt.Fprintln(w, "cs=", cs)
}

func main() {
	http.HandleFunc("/setCookie", setCookieHandle)
	http.HandleFunc("/getCookie", getCookieHandle)

	log.Println("listening...")
	http.ListenAndServe("127.0.0.1:9999", nil)
}
