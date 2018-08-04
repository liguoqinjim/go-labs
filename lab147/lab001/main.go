package main

import (
	"github.com/gocarina/gocsv"
	"os"
	"log"
)

type Client struct {
	// Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func main() {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("os.OpenFile error:%v", err)
	}
	defer clientsFile.Close()

	var clients []*Client
	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		log.Fatalf("goscv.UnmarshalFile error:%v", err)
	}
	for _, client := range clients {
		log.Printf("Id[%s],Name[%s],Age[%s]", client.Id, client.Name, client.Age)
	}
}
