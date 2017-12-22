package main

import (
	"bytes"
	"compress/flate"
	"io/ioutil"
	"log"
)

func main() {
	content := "A long time ago in a galaxy far, far away..."

	//压缩(用deflate)
	var buf bytes.Buffer
	zw, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		log.Fatalf("flate.NewWrite error:%v", err)
	}
	_, err = zw.Write([]byte(content))
	if err != nil {
		log.Fatalf("zw.Write error:%v", err)
	}
	if err := zw.Close(); err != nil {
		log.Fatalf("zw.Close error:%v", err)
	}
	log.Println("origin length=", len(content))
	log.Println("deflate length=", buf.Len())

	//解压
	zr := flate.NewReader(&buf)
	out, err := ioutil.ReadAll(zr)
	if err != nil {
		log.Fatalf("ioutil.ReadAll error:%v", err)
	}
	if err := zr.Close(); err != nil {
		log.Fatalf("zr.Close error:%v", err)
	}
	log.Println("out length=", len(out))
	log.Println("out=", string(out))
}
