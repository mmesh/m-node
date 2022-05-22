package router

import (
	"net"

	"mmesh.dev/m-lib/pkg/xlog"
)

func (ni *networkInterface) checkInterfaceAddr(ip string) bool {
	if len(ip) == 0 {
		return false
	}

	ifaceName := ni.devName()

	ifc, err := net.InterfaceByName(ifaceName)
	if err != nil {
		xlog.Errorf("Unable to get network interface %s: %v", ifaceName, err)
		return false
	}

	addrs, err := ifc.Addrs()
	if err != nil {
		xlog.Errorf("Unable to get interface %s addrs: %v", ifaceName, err)
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
