package main

import (
	"github.com/satori/go.uuid"
	"log"
)

func main() {
	u1 := uuid.NewV4()
	log.Printf("UUIDv4: %s\n", u1)

	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		log.Printf("Something gone wrong: %v", err)
	}
	log.Printf("Successfully parsed: %s", u2)
}
