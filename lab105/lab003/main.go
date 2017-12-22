package main

import (
	"bytes"
	"compress/flate"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//dict 里面是要压缩的内容里面的一些频繁或者说多次出现的substring，这样压缩和解压的时候会用字典里面的值来代替
	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` + `<meta name="` + `" content="`

	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`

	var b bytes.Buffer
	zw, err := flate.NewWriterDict(&b, flate.DefaultCompression, []byte(dict))
	if err != nil {
		log.Fatalf("flate.NewWriteDict error:%v", err)
	}
	if _, err := io.Copy(zw, strings.NewReader(data)); err != nil {
		log.Fatalf("io.Copy error:%v", err)
	}
	if err := zw.Close(); err != nil {
		log.Fatalf("zw.Close error:%v", err)
	}
	log.Println("origin length=", len(data))
	log.Println("deflate length=", b.Len())

	//解压
	log.Println("Decompressed output using the dictionary:")
	zr := flate.NewReaderDict(bytes.NewReader(b.Bytes()), []byte(dict))
	if out, err := ioutil.ReadAll(zr); err != nil {
		log.Fatalf("io.Copy error:%v", err)
	} else {
		log.Println("out1=\n", string(out))
	}
	if err := zr.Close(); err != nil {
		log.Fatalf("zr.Close error:%v", err)
	}

	//故意替换掉dict的值，查看解压之后的效果
	log.Println("Substrings matched by the dictionary are marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	zr = flate.NewReaderDict(&b, hashDict)
	if out, err := ioutil.ReadAll(zr); err != nil {
		log.Fatalf("io.Copy error:%v", err)
	} else {
		log.Println("out2=\n", string(out))
	}
	if err := zr.Close(); err != nil {
		log.Fatalf("zr.Close error:%v", err)
	}

}
