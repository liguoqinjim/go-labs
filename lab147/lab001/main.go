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
		log.Println("Hello", client.Name)
	}

	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		log.Fatalf("clientsFile.Seek error:%v", err)
	}

	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "Fred"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})
	csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	if err != nil {
		log.Fatalf("gocsv.MarshalString error:%v", err)
	}
	log.Println(csvContent) // Display all clients as CSV string

	err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		log.Fatalf("gocsv.MarshalFile error:%v", err)
	}
}
