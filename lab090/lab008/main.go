package main

import (
	"bytes"
	"compress/gzip"
	"github.com/parnurzeal/gorequest"
	"io"
	"log"
)

func main() {
	request := gorequest.New()

	_, body, errs := request.Get("http://httpbin.org/gzip").
		Set("Accept-Encoding", "gzip, deflate, br").
		End()
	if errs != nil {
		log.Fatalf("request.Get error:%v", errs)
	}
	log.Println("body=", body)

	newBody, err := gUnzipData([]byte(body))
	if err != nil {
		log.Fatalf("gUnzipData error:%v", err)
	}

	log.Println("newBody=", string(newBody))
}

func gUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}
