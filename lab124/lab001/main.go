package main

import (
	"github.com/juju/persistent-cookiejar"
	"log"
)

func main() {
	//查看默认保存位置
	filePath := cookiejar.DefaultCookieFile()
	log.Println("filePath=", filePath)
}
