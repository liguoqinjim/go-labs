package main

import (
	"github.com/huichen/sego"
	"log"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("../dictionary.txt")

	// 分词
	text := []byte("中华人民共和国中央人民政府")
	segments := segmenter.Segment(text)

	// 普通模式
	log.Println(sego.SegmentsToString(segments, false))

	// 搜索模式
	log.Println(sego.SegmentsToString(segments, true))
}
