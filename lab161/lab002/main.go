package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"log"
)

const (
	t = "This page intentionally left blank."
)

func main() {
	data := md5.Sum([]byte(t))
	log.Printf("%s", data)

	data_base64 := base64.URLEncoding.EncodeToString(data[:])
	log.Println(data_base64)

	data_hex := hex.EncodeToString(data[:])
	log.Println(data_hex)
}
