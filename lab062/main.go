package main

import (
	"fmt"
)

func main() {
	a := "公会名称1" //我们认为这个字符串的长度是5，中文一个字长度就是1，数字长度也是1
	fmt.Println("a=", a)
	fmt.Println("len(a)=", len(a)) //output 但是这里的长度是13,因为string底层是用byte来存的，一个中文是3个byte

	fmt.Println("len([]rune(a))=", len([]rune(a)))
}
