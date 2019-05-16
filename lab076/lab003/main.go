package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	//第二个
	log.Println("------------------------")
	const sample2 = "\x59\x69\x74\x74\x77\x35\x44\x44\x6e\x57\x6e\x44\x72\x41\x3d\x3d"

	log.Println("Println:")
	log.Println(sample2)

	var sample3 string
	sample3 = "\x59"
	log.Println(sample3)

	sample4 := sample3
	log.Println(sample4)

	sample5, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}
	sample6 := string(sample5)
	log.Println(sample6)
}
