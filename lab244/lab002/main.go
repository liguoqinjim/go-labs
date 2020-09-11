package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/tdewolff/minify/v2/xml"
)

func main() {
	j := `<x a="b"></x>`
	s := bytes.NewBufferString("")

	if err := xml.Minify(nil, s, strings.NewReader(j), nil); err != nil {
		log.Fatalf("xml.Minify error:%v", err)
	} else {
		log.Printf("j=%s", j)
		log.Printf("s=%s", s)
	}
}
