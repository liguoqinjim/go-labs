package main

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"
)

func main() {
	//压缩
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)

	content := "A long time ago in a galaxy far, far away..."
	_, err := zw.Write([]byte(content))
	if err != nil {
		log.Fatalf("zw.Write error:%v", err)
	}

	if err := zw.Close(); err != nil {
		log.Fatalf("zw.Close error:%v", err)
	}

	log.Println("origin length=", len(content))
	log.Println("zip length=", len(buf.Bytes()))

	//解压
	zr, err := zlib.NewReader(&buf)
	if err != nil {
		log.Fatalf("zlib.NewReader error:%v", err)
	}
	out, err := ioutil.ReadAll(zr)
	if err != nil {
		log.Fatalf("ioutil.ReadAll error:%v", err)
	}
	log.Println("unzip length=", len(out))

	if err := zr.Close(); err != nil {
		log.Fatalf("zr.Close error:%v", err)
	}
}
