package main

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
)

func GB2312ToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}

func UTF8ToGB2312(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {
	s := "这是一段测试"

	gb2312, err := UTF8ToGB2312([]byte(s))
	if err != nil {
		log.Printf("UTF8ToGB2312 error:%v", err)
	} else {
		log.Println("gb2312=", string(gb2312))
	}

	utf8, err := GB2312ToUTF8(gb2312)
	if err != nil {
		log.Printf("GB2312ToUTF8 error:%v", err)
	} else {
		log.Println("utf8=", string(utf8))
	}
}
