package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"io/ioutil"
	"log"
)

func main() {
	//生成到byte
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		log.Fatalf("qrcode.Encode error:%v", err)
	}
	if err := ioutil.WriteFile("1.png", png, 0600); err != nil {
		log.Fatalf("writeFile error:%v", err)
	}

	//生成文件
	if err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "2.png"); err != nil {
		log.Fatalf("writeFile error:%v", err)
	}

	//生成文件带颜色
	if err := qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.Black, color.White, "3.png"); err != nil {
		log.Fatalf("writeColorFile error:%v", err)
	}
}
