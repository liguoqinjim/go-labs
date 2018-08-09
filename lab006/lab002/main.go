package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//open a file
	f, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer f.Close()

	//读取5个字节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	if err != nil {
		log.Fatalf("f.Read error:%v", err)
	}
	log.Printf("%d bytes: %s\n", n1, string(b1))
	//Output 5 bytes: [{"Na

	//偏移位置，但是Read方法也会把游标移动到读取完的位置
	o2, err := f.Seek(6, 0) //whence为0的时候，表示从文件头计算偏移
	if err != nil {
		log.Fatalf("f.Seek error:%v", err)
	}
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	if err != nil {
		log.Fatalf("f.Read error:%v", err)
	}
	log.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//读取整个文件，每次读一个buf的大小，读到EOF停止
	//归位到0
	f.Seek(0, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			log.Print(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("f.Read error:%v", err)
		}
	}
}
