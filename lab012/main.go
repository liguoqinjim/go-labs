package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	n, err := w.WriteString("Hello go!\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %d bytes\n", n)

	w.Flush()
}
