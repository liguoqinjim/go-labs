package main

import (
	"fmt"
	"math/big"
	"net"
)

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func main() {
	//3055975544

	ip := "192.168.78.123"
	ipInt := InetAtoN(ip)

	fmt.Printf("convert string ip [%s] to int: %d\n", ip, ipInt)
	fmt.Printf("convert int ip [%d] to string: %s\n", ipInt, InetNtoA(ipInt))

	fmt.Println(InetNtoA(3055975544))
}
