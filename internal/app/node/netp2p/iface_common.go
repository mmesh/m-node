package netp2p

import (
	"net"

	"github.com/songgao/water"
	"github.com/spf13/viper"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

const MTU int = 1400 // TUN interface, so only plain IP packet, no ethernet header + mtu is set to 1300

var iface *water.Interface

func checkInterfaceAddr(ip string) bool {
	ifaceName := viper.GetString("agent.iface")

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
			//xlog.Warnf("Address %s is already configured on interface %s", ip, ifaceName)
			return false
		}
	}

	return true
}

func getConnectedRoutes() ([]*net.IPNet, error) {
	var connectedRoutes []*net.IPNet

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function net.InterfaceAddrs()", errors.Trace())
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok {
			connectedRoutes = append(connectedRoutes, ipNet)
		}
	}

	return connectedRoutes, nil
}
