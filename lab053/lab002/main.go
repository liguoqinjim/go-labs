package main

import (
	"github.com/hoisie/web"
	"log"
)

func hello(val string) string {
	return "hello " + val
}

func error2(err error) {
	log.Println("error!")
}

func main() {
	web.Get("/hello(.*)", hello)
	web.Get("/error", error2)

	web.Run("0.0.0.0:9999")
}
