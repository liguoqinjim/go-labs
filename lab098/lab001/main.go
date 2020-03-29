package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatalf("create file error:%v", err)
	}
	defer f.Close()

	// 写入UTF-8 BOM 解决excel里面的乱码
	if _, err := f.WriteString("\xEF\xBB\xBF"); err != nil {
		log.Fatalf("writeString error:%v", err)
	}

	//writeAll
	w := csv.NewWriter(f)
	datas := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
	}
	if err := w.WriteAll(datas); err != nil {
		log.Fatalf("w.WriteAll error:%v", err)
	}

	//write
	data := []string{
		"3", "日本", "23",
	}
	if err := w.Write(data); err != nil {
		log.Fatalf("w.Write error:%v", err)
	}

	w.Flush()
}
