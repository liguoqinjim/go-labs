package main

import (
	"fmt"
	"github.com/anvie/port-scanner"
	"time"
)

func main() {
	//for i := 211; i <= 221; i++ {

	for i := 17; i <= 20; i++ {
		host := fmt.Sprintf("15.17.26.%d", i)

		ps := portscanner.NewPortScanner(host, 2*time.Second, 5)

		fmt.Printf("scanning %s port %d-%d...\n", host, 7400, 7600)

		openedPorts := ps.GetOpenedPort(20, 65535)

		for i := 0; i < len(openedPorts); i++ {
			port := openedPorts[i]
			fmt.Print(" ", port, " [open]")
			fmt.Println("  -->  ", ps.DescribePort(port))
		}
	}
}
