package main

import (
	"github.com/parnurzeal/gorequest"
	"io"
	"log"
	"os"
)

func main() {
	request := gorequest.New()

	resp, _, errs := request.Get("http://weixin.sogou.com/antispider/util/seccode.php").
		End()
	if errs != nil {
		log.Fatalf("errs:%v", errs)
	}

	fi, err := os.Create("captcha.jpg")
	if err != nil {
		log.Fatalf("os.Create error:%v", err)
	}
	defer fi.Close()

	_, err = io.Copy(fi, resp.Body)
	if err != nil {
		log.Fatalf("")
	}
	defer resp.Body.Close()
}
