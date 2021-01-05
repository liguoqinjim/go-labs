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
	send(port, []byte("AT+CMGS=\"13888888888\"\r\n"))

	sms := []byte("hello world123")
	sms = append(sms, 0x1a)
	send(port, sms)
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
