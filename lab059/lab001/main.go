package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs() //返回本机所有的地址
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range addrs {
		fmt.Println(v)
	}

	//整理地址
	fmt.Println()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() { //去除回环地址
			if ipnet.IP.To4() != nil { //去除ipv6
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
