package main

import (
	"github.com/dutchcoders/go-ouitools"
	"log"
)

func main() {
	db := ouidb.New("oui_database.txt")
	//db := ouidb.New("oui2.txt")
	if db == nil {
		log.Fatalf("db is nil")
	}

	mac := "D0-C7-C0-9A-EC-C8" //D0-C7-C0-9A-EC-C8
	v, err := db.VendorLookup(mac)
	if err != nil {
		log.Fatalf("parse: %s: %s", mac, err.Error())
	}

	log.Printf("%s => %s\n", mac, v)
}
