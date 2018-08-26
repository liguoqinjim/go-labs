package main

import (
	"encoding/xml"
	"log"
)

type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"config"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Addresses []Address `xml:"address"`
	Comment   string    `xml:",comment"`
}

func main() {
	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Addresses = make([]Address, 2)
	v.Addresses[0] = Address{City: "Hanga Roa", State: "Easter Island"}
	v.Addresses[1] = Address{City: "Hanga Roa2", State: "Easter Island2"}

	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("xml.MarshalIndent error:%v", err)
	}

	log.Printf("\n%s", output)
}
