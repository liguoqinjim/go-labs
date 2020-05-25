package main

import (
	"io/ioutil"
	"log"

	"github.com/h2non/filetype"
)

func main() {
	paths := []string{"../data/1.png", "../data/2.jpeg", "../data/3.png"}

	for _, path := range paths {
		getType(path)
	}
}

func getType(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	kind, err := filetype.Match(buf)
	if err != nil {
		log.Fatalf("filetype.Match error:%v", err)
	}
	if kind == filetype.Unknown {
		log.Println("Unknown file type")
		return
	}

	log.Printf("File type: %s. MIME: %s", kind.Extension, kind.MIME.Value)
}
