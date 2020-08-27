package main

import (
	"net"
	"regexp"
	"strings"
)

func IsIpv4Net(host string) bool {
	return net.ParseIP(host) != nil
}

var (
	ipRegex, _ = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
)

func IsIpv4Regex(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	return ipRegex.MatchString(ipAddress)
}

const (
	IPv4len = 4
	big     = 255
)

func IsIpv4(s string) bool {
	for i := 0; i < IPv4len; i++ {
		if len(s) == 0 {
			return false
		}

		if i > 0 {
			if s[0] != '.' {
				return false
			}
			s = s[1:]
		}

		var n int
		var j int
		for j = 0; j < len(s) && '0' <= s[j] && s[j] <= '9'; j++ {
			n = n*10 + int(s[j]-'0')
		}
		if j == 0 {
			n = 0
			continue
		}
		if n > 0xFF {
			return false
		}
		s = s[j:]
	}

	if len(s) != 0 {
		return false
	}
	return true
}
