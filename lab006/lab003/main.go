package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("test.json")
	checkError(err)
	defer f.Close()

	buf := bufio.NewReaderSize(f, 0)

	for {
		line, err := buf.ReadBytes('\n')
		fmt.Printf("%s", line)

		if err == io.EOF {
			break
		}
	}
}
