package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Encoding to use. Since this implements the encoding.Encoding
// interface from golang.org/x/text/encoding you can trivially
// change this out for any of the other implemented encoders,
// e.g. `traditionalchinese.Big5`, `charmap.Windows1252`,
// `korean.EUCKR`, etc.
var enc = simplifiedchinese.GBK

func main() {
	const filename = "data/1.html"

	//exampleWriteGBK(filename)
	//exampleReadGBK(filename)

	//exampleRead(filename)
	exampleReadGBK01(filename)
}

//直接读会有乱码
func exampleRead(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(data)
	}

	fmt.Printf("%s", data)
}

func exampleReadGBK01(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(data)
	}

	//data里面是gbk，所有用decoder
	reader := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s", d)
}

func exampleReadGBK02(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := transform.NewReader(f, enc.NewDecoder())

	// Read converted UTF-8 from `r` as needed.
	// As an example we'll read line-by-line showing what was read:
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Printf("Read line: %s\n", sc.Bytes())
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}
}
