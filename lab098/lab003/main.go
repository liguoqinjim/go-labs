package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatalf("create file error:%v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	//data
	data := []map[string]interface{}{
		{"name": "tom", "age": 12},
		{"name": "jack", "age": 14},
	}
	headers := []string{"name", "age"}

	//header
	if err := w.Write(headers); err != nil {
		log.Fatalf("w.Write error:%v", err)
	}

	//write
	for _, d := range data {
		ds := make([]string, len(d))

		for n, v := range headers {
			ds[n] = fmt.Sprintf("%v", d[v])
		}

		if err := w.Write(ds); err != nil {
			log.Fatalf("w.Write error:%v", err)
		}
	}

	w.Flush()
}
