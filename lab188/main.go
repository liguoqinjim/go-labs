package main

import (
	"github.com/dutchcoders/go-ouitools"
	"log"
)

func main() {
	db := ouidb.New("oui.txt")
	if db == nil {
		log.Fatalf("db is nil")
	}

	mac:="00:16:e0:3d:f4:4c"
	v, err := db.VendorLookup(mac)
	if err != nil {
		log.Fatalf("parse: %s: %s", mac, err.Error())
	}

	log.Printf("%s => %s\n", mac, v)
}
