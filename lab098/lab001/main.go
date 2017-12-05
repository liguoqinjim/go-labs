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

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	datas := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
	}
	w.WriteAll(datas)

	data := []string{
		"3", "日本", "23",
	}
	w.Write(data)

	w.Flush()
}
