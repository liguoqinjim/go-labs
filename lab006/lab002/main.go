package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//open a file
	f, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	//close a file
	defer f.Close()

	//读取5个字节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//偏移位置，但是Read方法也会把游标移动到读取完的位置
	o2, err := f.Seek(6, 0)
	checkError(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	checkError(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//归位到0
	f.Seek(0, 0)

	//读取整个文件，每次读一个buf的大小，读到EOF停止
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
		checkError(err)
	}
}
