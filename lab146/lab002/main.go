package main

import (
	"github.com/jmz331/gpinyin"
	"log"
)

func main() {
	//有声调
	const s = "你好世界"
	r := gpinyin.ConvertToPinyinString(s, "-", gpinyin.PINYIN_WITH_TONE_MARK)
	log.Println(r)

	//没有声调
	r2 := gpinyin.ConvertToPinyinString(s, "-", gpinyin.PINYIN_WITHOUT_TONE)
	log.Println(r2)

	//要是有不是中文的字符会跳过
	const s2 = "你好adc 世界?"
	r3 := gpinyin.ConvertToPinyinString(s2, "-", gpinyin.PINYIN_WITHOUT_TONE)
	log.Println(r3)
}
