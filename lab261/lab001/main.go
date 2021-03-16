package main

import (
	"bytes"
	"github.com/dhowden/tag"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	demo()
}

func demo() {
	f, err := os.Open("../data/output.mp3")
	//f, err := os.Open("../../lab262/data/test.mp3")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).
	log.Println(m.Album())
	albumUtf, _ := GBKToUTF8([]byte(m.Album()))
	log.Printf("%s", albumUtf)

	log.Println(m.Artist())
	artistUtf, _ := GBKToUTF8([]byte(m.Artist()))
	log.Printf("%s", artistUtf)
}

func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}
