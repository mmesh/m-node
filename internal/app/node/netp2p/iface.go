package netp2p

import (
	"net"

	"mmesh.dev/m-lib/pkg/xlog"
)

func checkInterfaceAddr(ip string) bool {
	if iface == nil || len(ip) == 0 {
		return false
	}

	ifc, err := net.InterfaceByName(iface.name)
	if err != nil {
		xlog.Errorf("Unable to get network interface %s: %v", iface.name, err)
		return false
	}

	addrs, err := ifc.Addrs()
	if err != nil {
		xlog.Errorf("Unable to get interface %s addrs: %v", iface.name, err)
		return false
	}

	for _, addr := range addrs {
		var ipAddr net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ipAddr = v.IP
		case *net.IPAddr:
			ipAddr = v.IP
		}

		if ipAddr.Equal(net.ParseIP(ip)) {
			// xlog.Warnf("Address %s is already configured on interface %s", ip, iface.name)
			return false
		}
	}

	return true
}
