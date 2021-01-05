package main

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:              "/dev/cu.usbserial-FTA3JNIA",
		BaudRate:              9600,
		DataBits:              8,
		StopBits:              1,
		ParityMode:            serial.PARITY_NONE,
		RTSCTSFlowControl:     false,
		InterCharacterTimeout: 200,
		//MinimumReadSize:   2,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	} else {
		log.Println("serial.Open success")
	}

	// Make sure to close it later.
	defer port.Close()

	send(port, []byte("at+ipr?\r\n"))
	send(port, []byte("AT+CMGF=1\r\n"))
	send(port, []byte("AT+csmp=17,167,1,8\r\n"))
	send(port, []byte("AT+cscs=\"ucs2\"\r\n"))

	send(port, []byte("AT+CMGS=\"00310033003800380038003800380038003800380038\"\r\n")) //13888888888

	data := []byte("6B228FCE4F7F752898DE601D521B00530049004D00380030003000436A215757FF01")
	data = []byte("4F60597D4E16754C") //你好世界
	data = append(data, 0x1a)
	send(port, data)
}

func send(port io.ReadWriteCloser, data []byte) {
	n, err := port.Write(data)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	} else {
		log.Println("port.Write", n, "bytes.")
	}

	res, err := ioutil.ReadAll(port)
	if err != nil {
		log.Fatalf("read error:%v", err)
	} else {
		log.Printf("response:%s", res)
	}
}
