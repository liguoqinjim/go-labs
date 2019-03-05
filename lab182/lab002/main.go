package main

import (
	"fmt"
	"github.com/anvie/port-scanner"
	"time"
)

func main() {
	for i := 211; i <= 221; i++ {

	}

	ps := portscanner.NewPortScanner("hp-111", 2*time.Second, 5)

	fmt.Printf("scanning port %d-%d...\n", 20, 65535)

	openedPorts := ps.GetOpenedPort(20, 30000)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}
}
