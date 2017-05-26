package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	//本机ip
	ips, err := getIntranetIP()
	LogError(err)

	if len(ips) != 1 {
		log.Fatal("得到本机ip error length > 1")
	}

	for _, v := range ips {
		fmt.Println("本机ip", v)
	}

	//公网ip
	ip, err := getInternetIP()
	LogError(err)
	fmt.Println("公网ip", ip)
}

func getIntranetIP() ([]string, error) { //得到本机ip地址
	ifaces, err := net.Interfaces()
	LogError(err)

	ips := make([]string, 0)

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 { //状态不是up
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 { //去掉回环地址(==1的时候是回环地址)
			continue
		}

		if strings.HasPrefix(iface.Name, "VM") || strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") { //去除虚拟网卡
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return nil, err
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

			ip = ip.To4()
			if ip == nil { //去除ipv4
				continue
			}

			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}

	return ips, nil
}

func IsIntranet(ip string) bool { //是否是内网地址
	//A类 10.0.0.0--10.255.255.255
	//B类 172.16.0.0--172.31.255.255  B类要特殊处理，不像A类和C类只要判断是否有前缀就可以
	//C类 192.168.0.0--192.168.255.255

	if strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "192.168.") {
		return true
	}

	if strings.HasPrefix(ip, "172.") {
		arr := strings.Split(ip, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.Atoi(arr[1])
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

func getInternetIP() (string, error) { //得到公网ip (http://myexternalip.com/raw)
	resp, err := http.Get("http://myexternalip.com/raw")
	LogError(err)

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	LogError(err)

	exip := string(bytes.TrimSpace(b))
	return exip, nil
}

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
