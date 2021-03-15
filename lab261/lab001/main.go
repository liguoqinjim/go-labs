package main

import (
	"github.com/dhowden/tag"
	"log"
	"os"
)

func main() {
	demo()
}

func demo() {
	f, err := os.Open("../data/output.mp3")
	if err != nil {
		log.Fatalf("os.Open error:%v", err)
	}

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).
	log.Println(m.Album())
}
