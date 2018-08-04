package main

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

type Client struct {
	// Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func main() {
	var clients []*Client
	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "Fred"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})

	//生成csv string
	csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	if err != nil {
		log.Fatalf("gocsv.MarshalString error:%v", err)
	}
	log.Println("\n" + csvContent) // Display all clients as CSV string

	//生成csv文件
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("os.OpenFile error:%v", err)
	}
	defer clientsFile.Close()
	err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		log.Fatalf("gocsv.MarshalFile error:%v", err)
	}
}
