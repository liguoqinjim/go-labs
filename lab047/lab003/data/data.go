package data

import (
	"io"
	"os"
)

var BattleData string

func LoadData() {
	file, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	BattleData = string(chunks)
}
