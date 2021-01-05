package main

import (
	"github.com/jacobsa/go-serial/serial"
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

	// Write 4 bytes to the port.
	n, err := port.Write(append([]byte("AT+IPR?\r\n")))
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	} else {
		log.Println("port.Write", n, "bytes.")
	}

	data, err := ioutil.ReadAll(port)
	if err != nil {
		log.Fatalf("read error:%v", err)
	} else {
		log.Printf("response:%s", data)
	}

	//buf := make([]byte, 32)
	//n, err = port.Read(buf)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Printf("%q", buf[:n])
	//}
}
