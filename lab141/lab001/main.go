package main

import (
	"lab141/lab001/web"
	"log"
	"net/http"
)

func main() {
	web.Router()

	log.Println("Listening...")
	http.ListenAndServe(":9999", nil)
}
