package main

import (
	"github.com/teris-io/shortid"
	"log"
)

func main() {
	for i := 0; i < 10; i++ {
		log.Println(shortid.Generate())
	}

	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatalf("shortid.New error:%v", err)
	}

	// then either:
	log.Println(sid.Generate())
	log.Println(sid.Generate())

	// or:
	shortid.SetDefault(sid)
	// followed by:
	log.Println(shortid.Generate())
	log.Println(shortid.Generate())
}
