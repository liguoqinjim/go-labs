package main

import (
	"bytes"
	"github.com/frolovo22/tag"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
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
	//f, err := os.Open("../data/output_gbk_origin.mp3")
	f, err := os.Open("../data/modify.mp3")
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

	log.Println("artist", artist)
	return

	artistUtf, err := GBKToUTF8([]byte(artist))
	if err != nil {
		log.Fatalf("GBKToUTF8 error:%v", err)
	}
	artistUtf = []byte("你好")

	artistUtf16, err := UTF8To16(artistUtf)
	if err != nil {
		log.Fatalf("utf8 to 16 error:%v", err)
	}
	log.Printf("artistUtf16:%s", artistUtf16)

	if err := tags.DeleteArtist(); err != nil {
		log.Fatalf("delete artist error:%v", err)
	}

	if err := tags.SetArtist(string(artistUtf16)); err != nil {
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

func UTF8To16(s []byte) ([]byte, error) {
	decoder := unicode.UTF8.NewDecoder()
	return decoder.Bytes(s)
}
