package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	var str = "hello world"

	hasher := md5.New()
	hasher.Write([]byte(str))
	fmt.Println(str)
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}
