package utils

import "net"

func IPv4IsValid(ipv4 string) bool {
	if len(ipv4) == 0 {
		return false
	}

	ip := net.ParseIP(ipv4)

	if ip == nil {
		return false
	}

	if ip.IsLoopback() || ip.IsMulticast() {
		return false
	}

	if !ip.IsGlobalUnicast() {
		return false
	}

	return true
}
