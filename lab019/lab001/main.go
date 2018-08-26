package main

import (
	"encoding/xml"
	"log"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address struct {
	City, State string
}
type Result struct {
	XMLName xml.Name `xml:"config"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}

func main() {
	v := Result{Name: "none", Phone: "none"}

	data := `
		<config>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</config>
	`

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %v", err)
	}

	log.Printf("XMLName: %#v\n", v.XMLName)
	log.Printf("Name: %q\n", v.Name)
	log.Printf("Phone: %q\n", v.Phone)
	log.Printf("Email: %v\n", v.Email)
	log.Printf("Groups: %v\n", v.Groups)
	log.Printf("Address: %v\n", v.Address)
}
