package main

import (
	"bytes"
	"github.com/frolovo22/tag"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	demo()
}

//修改标签
func demo() {
	f, err := os.Open("../data/output_gbk_origin.mp3")
	if err != nil {
		log.Fatalf("os open error:%v", err)
	}

	tags, err := tag.Read(f)
	if err != nil {
		log.Fatalf("tag.ReadFile error:%v", err)
	}
	log.Println(tags.GetTitle())

	tagVersion := tags.GetVersion()
	log.Println("tagVersion", tagVersion)

	//album
	album, err := tags.GetAlbum()
	if err != nil {
		log.Fatalf("get album error:%v", err)
	}
	albumUtf, err := GBKToUTF8([]byte(album))
	if err != nil {
		log.Fatalf("gbk to utf error:%v", err)
	}
	log.Printf("album utf:%s", albumUtf)

	albumUtf = []byte("nihao")
	if err := tags.SetAlbum(string(albumUtf)); err != nil {
		log.Fatalf("set album error:%v", err)
	}

	//artist
	artist, err := tags.GetArtist()
	if err != nil {
		log.Fatalf("get artist error:%v", err)
	}
	artistUtf, err := GBKToUTF8([]byte(artist))
	if err != nil {
		log.Fatalf("GBKToUTF8 error:%v", err)
	}
	artistUtf = []byte("hello")
	if err := tags.SetArtist(string(artistUtf)); err != nil {
		log.Fatalf("set artist error:%v", err)
	}

	if err := tags.SaveFile("../data/modify.mp3"); err != nil {
		log.Fatalf("save error:%v", err)
	}
}

func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}
