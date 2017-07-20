package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	//第一种方法： EncodeToString,DecodeString
	fmt.Println("第一种方法")
	//string->hex
	data := "Hello Gopher!"
	fmt.Println("data:", data)
	src := []byte(data)
	encodedStr := hex.EncodeToString(src)
	fmt.Println("string->hex:", encodedStr)

	//hex->string
	decoded, err := hex.DecodeString(encodedStr)
	if err != nil {
		log.Fatal("DecodeString error:", err)
	}
	fmt.Println("hex->string:", string(decoded))

	//第二种方法：Encode,Decode
	fmt.Println("\n第二种方法")
	//string->hex
	src2 := []byte("Hello Gopher!")
	dst2 := make([]byte, hex.EncodedLen(len(src2)))
	hex.Encode(dst2, src2)
	fmt.Printf("string->hex: %s\n", dst2)

	//hex->string
	dst3 := make([]byte, hex.DecodedLen(len(dst2)))
	n, err := hex.Decode(dst3, dst2)
	if err != nil {
		log.Fatal("Decode error:", err)
	}

	fmt.Printf("hex->string %s\n", dst3[:n])
}
