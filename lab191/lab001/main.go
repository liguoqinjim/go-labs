package main

import (
	"encoding/xml"
	"log"

	"github.com/tiaguinho/gosoap"
)

// GetIPLocationResponse will hold the Soap response
type GetIPLocationResponse struct {
	GetIPLocationResult string `xml:"GetIpLocationResult"`
}

// GetIPLocationResult will
type GetIPLocationResult struct {
	XMLName xml.Name `xml:"GeoIP"`
	Country string   `xml:"Country"`
	State   string   `xml:"State"`
}

var (
	r GetIPLocationResponse
)

func main() {
	soap, err := gosoap.SoapClient("http://wsgeoip.lavasoft.com/ipservice.asmx?WSDL")
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.Params{
		"sIp": "8.8.8.8",
	}

	err = soap.Call("GetIpLocation", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

	xml.Unmarshal(soap.Body, &r)

	// GetIpLocationResult will be a string. We need to parse it to XML
	result := GetIPLocationResult{}
	err = xml.Unmarshal([]byte(r.GetIPLocationResult), &result)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	if result.Country != "US" {
		log.Fatalf("error: %+v", r)
	}

	log.Println("Country: ", result.Country)
	log.Println("State: ", result.State)
}
