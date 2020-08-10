package main

import (
	"log"
	"strconv"
)

func main() {
	toASCII()
	fromACSII()
}

//啊->\u554a
func toASCII() {
	quoted := strconv.QuoteRuneToASCII('啊') // quoted = "'\u554a'"
	log.Println("quoted:", quoted)
	unquoted := quoted[1 : len(quoted)-1] // unquoted = "\u554a"
	log.Println("unquoted:", unquoted)

	quoted = strconv.QuoteToASCII("你好啊")
	log.Println("quoted:", quoted)
	unquoted = quoted[1 : len(quoted)-1]
	log.Println("unquoted:", unquoted)
}

//  \u4f60\u597d\u554a->你好啊
func fromACSII() {
	unquote, err := strconv.Unquote(`"\u4f60\u597d\u554a"`)
	if err != nil {
		log.Fatalf("Unquote error:%v", err)
	}
	log.Println("unquote:", unquote)

	unquote, err = strconv.Unquote(`\u4f60\u597d\u554a`)
	if err != nil {
		log.Fatalf("Unquote error:%v", err)
	}
	log.Println("unquote:", unquote)
}
