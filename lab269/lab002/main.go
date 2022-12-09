package main

import (
	"github.com/mojocn/base64Captcha"
	"log"
)

func main() {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.3, 45)
	driver = base64Captcha.DefaultDriverDigit

	item, err := driver.DrawCaptcha("12345")
	if err != nil {
		log.Fatalf("driver.DrawCaptcha error:%v", err)
	}

	log.Println(item.EncodeB64string())
	//https://jaredwinick.github.io/base64-image-viewer/ 查看
}
