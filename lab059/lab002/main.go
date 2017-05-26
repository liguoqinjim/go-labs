package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	ips, err := IntranetIP()
	if err != nil {
		log.Fatal(ips)
	}

	for _, v := range ips {
		fmt.Println(v)
	}
}

func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)

	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}

	for _, v := range ifaces {
		fmt.Println(v)
	}
	fmt.Println()

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 { //网卡是否连接
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 { //是否是回环地址
			continue // loopback interface
		}

		// ignore docker and warden bridge 去除docker网卡
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4() //去除ipv6
			if ip == nil {
				continue // not an ipv4 address
			}

			ipStr := ip.String()
			if IsIntranet(ipStr) { //判断是否是内网地址
				ips = append(ips, ipStr)
			}
		}
	}

	return ips, nil
}

func IsIntranet(ipStr string) bool { //是否是内网地址
	//A类 10.0.0.0--10.255.255.255
	//B类 172.16.0.0--172.31.255.255  B类要特殊处理，不像A类和C类只要判断是否有前缀就可以
	//C类 192.168.0.0--192.168.255.255

	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}

	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}
