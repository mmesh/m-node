package maddr

import (
	"fmt"
	"net"
	"strings"

	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func GetLocalMAddrs(port int) ([]multiaddr.Multiaddr, error) {
	maddrs := make([]multiaddr.Multiaddr, 0)

	agentIfaceName := viper.GetString("agent.iface")
	if len(agentIfaceName) == 0 {
		return nil, fmt.Errorf("missing agent.iface configuration")
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function net.Interfaces()", errors.Trace())
	}

	for _, ifc := range ifaces {
		if ifc.Name == agentIfaceName || ifc.Name == "lo" {
			continue
		}

		addrs, err := ifc.Addrs()
		if err != nil {
			xlog.Errorf("Unable to get interface %s addrs: %v", ifc.Name, err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if strings.Contains(ip.String(), ":") {
				// ipv6 not supported
				continue
			}

			if !ip.IsGlobalUnicast() {
				continue
			}

			var maTCP, maQUIC multiaddr.Multiaddr
			var err error

			// maddr for tcp transport
			maTCP, err = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", ip.String(), port))
			if err != nil {
				return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr()", errors.Trace())
			}
			maddrs = append(maddrs, maTCP)

			// maddr for udp quic transport
			maQUIC, err = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/%d/quic", ip.String(), port))
			if err != nil {
				return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr()", errors.Trace())
			}
			maddrs = append(maddrs, maQUIC)
		}
	}

	return maddrs, err
}

func String(maddrs ...multiaddr.Multiaddr) []string {
	s := make([]string, 0)

	for _, ma := range maddrs {
		s = append(s, ma.String())
	}

	return s
}
