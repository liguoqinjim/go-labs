package main

import (
	"github.com/frolovo22/tag"
	"io/ioutil"
	"log"
)

func main() {
	demo()
}

func demo() {
	frames := []tag.ID3v23Frame{
		{Key: "TPE1", Value: []byte("你好")},
		{Key: "TALB", Value: []byte("上海")},
	}

	data, err := ioutil.ReadFile("../data/output_gbk_origin.mp3")
	if err != nil {
		log.Fatalf("read file error:%v", err)
	}

	tags := tag.ID3v23{
		Frames: frames,
		Data:   data,
	}

	if err := tags.SaveFile("../data/lab002.mp3"); err != nil {
		log.Fatalf("save file error:%v", err)
	}
}
