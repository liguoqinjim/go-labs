package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
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
	s := "这是一段测试"
	gbk, err := UTF8ToGBK([]byte(s))
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("gbk=%s\n", gbk)
	}

	utf8, err := GBKToUTF8(gbk)
	if err != nil {
		fmt.Printf("GBKToUTF8 error:%v", err)
	} else {
		fmt.Printf("utf8=%s\n", utf8)
	}
}
