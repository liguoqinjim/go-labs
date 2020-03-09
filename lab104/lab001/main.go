package main

import (
	"log"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u1 := uuid.NewV4()
	log.Printf("UUID:%s", u1.String())

	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		log.Printf("Something gone wrong: %v", err)
	}
	log.Printf("Successfully parsed: %s", u2)
}
