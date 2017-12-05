package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatalf("open file error:%v", err)
	}

	reader := csv.NewReader(bufio.NewReader(f))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("read error:%v", err)
		}

		for _, v := range line {
			log.Println(v)
		}
	}
}
