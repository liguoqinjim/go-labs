package main

import (
	"bytes"
	"compress/gzip"
	"log"
	//"time"
	"io/ioutil"
)

func main() {
	//压缩
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	//setting the header fields is optional
	zw.Name = "1.txt"
	//zw.Comment = "an epic space opera by George Lucas"
	//zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	content := "A long time ago in a galaxy far, far away..."
	_, err := zw.Write([]byte(content))
	if err != nil {
		log.Fatalf("zw.write error:%v", err)
	}

	if err := zw.Close(); err != nil {
		log.Printf("zw.Close error:%v", err)
	}

	log.Println("origin length=", len(content))
	log.Println("gzip length=", buf.Len())

	//解压
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatalf("gzip.NewReader error:%v", err)
	}

	log.Printf("Name: %s", zr.Name)

	out, err := ioutil.ReadAll(zr)
	if err != nil {
		log.Fatalf("ioutils.ReadAll error:%v", err)
	}
	log.Println("gunzip length=", len(out))

	if err := zr.Close(); err != nil {
		log.Fatalf("zr.Close error:%v", err)
	}
}
