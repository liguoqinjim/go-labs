package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"log"
)

const (
	t = "This page intentionally left blank."
)

func main() {
	data := []byte(t)
	log.Printf("%s", sha1.Sum(data))

	hasher := sha1.New()
	hasher.Write([]byte(t))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	log.Println(sha)

	h := sha1.New()
	h.Write([]byte(t))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	log.Println(sha1_hash)
}
