package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		log.Fatalf("http.Get error:%v", err)
	}
	cookies := resp.Cookies()
	for _, c := range cookies {
		log.Println(c)
	}

	//cookies写入文件
	f, err := os.Create("cookies.txt")
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	for n, c := range cookies {
		_, err := w.WriteString(c.String())
		if err != nil {
			log.Fatalf("w.WriteString error")
		}

		if n != len(cookies)-1 {
			w.WriteString("\n")
		}
	}

	w.Flush()
	f.Close()

	//读取文件
	log.Println("读取文件")
	f, err = os.Open("cookies.txt")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}
	defer f.Close()

	req := &http.Request{}
	header := http.Header{}
	req.Header = header

	buf := bufio.NewReaderSize(f, 0)
	for {
		line, err := buf.ReadString('\n')
		log.Printf("%s", line)

		header.Add("Cookie", line)

		if err == io.EOF {
			break
		}
	}

	log.Println("cookies", req.Cookies())
}
