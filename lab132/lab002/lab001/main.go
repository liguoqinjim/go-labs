package main

import (
	"github.com/sparrc/go-ping"
	"log"
)

func main() {
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 2
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	log.Println("PacketsRecv=", stats.PacketsRecv)
	log.Println("PacketsSent=", stats.PacketsSent)
	log.Println("PacketLoss=", stats.PacketLoss)
	log.Println("Addr=", stats.Addr)
	log.Println("Rtts=", stats.Rtts)
}
