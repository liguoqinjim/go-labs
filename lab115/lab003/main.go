package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

var (
	salt = "f8d3d4e55c4de"
)

func main() {
	m5 := md5.New()
	m5.Write([]byte("mypassword"))
	m5.Write([]byte(salt))

	st := m5.Sum(nil)
	log.Printf("st=%s", st)
	stHex := hex.EncodeToString(st)
	log.Printf("st hex=%s", stHex)
}
