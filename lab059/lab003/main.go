package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	ip, err := getExternalIPAddr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ip=", ip)
}

func getExternalIPAddr() (exip string, err error) {
	resp, err := http.Get("http://myexternalip.com/raw") //通过这个网站来得到正确的ip
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	exip = string(bytes.TrimSpace(b))
	return
}
