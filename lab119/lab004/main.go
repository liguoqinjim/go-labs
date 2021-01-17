package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	digit()
	log.Println("--------------------------")

	toASCII()

	log.Println("--------------------------")
	toASCIIWithoutU()

	log.Println("--------------------------")
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

func toASCIIWithoutU() {
	quoted := strconv.QuoteToASCII("你好啊")
	log.Println("quoted:", quoted)
	unquoted := quoted[1 : len(quoted)-1]
	log.Println("unquoted:", unquoted)
	unquoted = strings.ReplaceAll(unquoted, "\\u", "")
	log.Println("unquoted without \\u:", unquoted)

	//upper
	unquoted = strings.ToUpper(unquoted)
	log.Println("unquoted upper:", unquoted)
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

//123456->003100320033003400350036
func digit() {
	ss := "123456"
	result := ""
	for _, s := range ss {
		result += fmt.Sprintf("00%x", s)
	}

	log.Println("result=", result)
}
