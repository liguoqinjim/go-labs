package main

import (
	"github.com/lithammer/shortuuid/v3"
	"log"
)

func main() {
	u := shortuuid.New() // Cekw67uyMpBGZLRP2HFVbe
	log.Println(u)

	u2 := shortuuid.NewWithNamespace("http://example.com")
	log.Println(u2)

	alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxy="
	u3 := shortuuid.NewWithAlphabet(alphabet) // u=BFWRLr5dXbeWf==iasZi
	log.Println(u3)
}
