package main

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
)

func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}

func UTF8ToGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {
	s := "GBK 与 UTF-8 编码转换测试"
	gbk, err := UTF8ToGBK([]byte(s))
	if err != nil {
		log.Printf("UTF8ToGBK error:%v", err)
	} else {
		log.Println("gbk=", string(gbk))
	}

	utf8, err := GBKToUTF8(gbk)
	if err != nil {
		log.Printf("GBKToUTF8 error:%v", err)
	} else {
		log.Println("utf8=", string(utf8))
	}
}
