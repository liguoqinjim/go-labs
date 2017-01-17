package main

import (
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

	d2 := []byte("Hello\ngo\n")
	n2, err := f.Write(d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("Hello goNuts")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
}
