package main

import (
	"log"
	"bufio"
	"strings"
	"io"
)

func main() {
	a := `this is the first line 
the second line
the third line`
	log.Printf("string=\n%s", a)

	//按行读取字符串
	r := bufio.NewReader(strings.NewReader(a))
	for {
		s, err := r.ReadString('\n')
		log.Print(s)
		if err == io.EOF {
			break
		}
	}
}
