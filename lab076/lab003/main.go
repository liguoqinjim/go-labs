package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	const sample = "\x66\x69\x72\x73\x74\x53\x74\x61\x74\x65"
	log.Println(sample)
	log.Println([]byte(sample))

	const sample2 = `\x66\x69\x72\x73\x74\x53\x74\x61\x74\x65`
	log.Println(sample2)
	log.Println([]byte(sample2))

	str := strconv.Quote(sample2)
	replaced := strings.Replace(str, `\\`, "\\", -1)
	newstr, err := strconv.Unquote(replaced)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(newstr)
}
