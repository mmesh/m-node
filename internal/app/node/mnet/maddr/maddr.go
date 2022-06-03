package maddr

import (
	"net"
	"strings"

	"github.com/multiformats/go-multiaddr"
)

func GetGlobalUnicastAddrStrings(maddrs ...multiaddr.Multiaddr) []string {
	s := make([]string, 0)

	for _, ma := range maddrs {
		sma := strings.Split(ma.String(), "/")

		if len(sma) < 5 {
			continue
		}

		ip := net.ParseIP(sma[2])
		if ip == nil {
			continue
		}

		if !ip.IsGlobalUnicast() {
			continue
		}

		s = append(s, ma.String())
	}

	return s
}
