package main

import (
	"bufio"
	"fmt"
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
	exampleReadGBK(filename)
}

func exampleReadGBK(filename string) {
	// Read UTF-8 from a GBK encoded file.
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