package netp2p

import (
	"fmt"
	"net"
	"strings"

	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/viper"
	"x6a.dev/pkg/xlog"
)

type addrList []multiaddr.Multiaddr

func (al *addrList) string() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *addrList) set(value string) error {
	addr, err := multiaddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}

func (al *addrList) strings() []string {
	var strs []string
	for _, ma := range *al {
		strs = append(strs, fmt.Sprintf("%v", ma))
	}

	return strs
}

func stringsToMAddrs(addrStrings []string) (maddrs []multiaddr.Multiaddr, err error) {
	for _, addrString := range addrStrings {
		addr, err := multiaddr.NewMultiaddr(addrString)
		if err != nil {
			return maddrs, err
		}
		maddrs = append(maddrs, addr)
	}
	return
}

func getLocalMAddrs(port int32) []multiaddr.Multiaddr {
	var maddrs []multiaddr.Multiaddr

	agentIfaceName := viper.GetString("agent.iface")

	ifaces, err := net.Interfaces()
	if err != nil {
		xlog.Errorf("Unable to get network interfaces: %v", err)
		return nil
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

			ma, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", ip.String(), port))
			if err != nil {
				xlog.Errorf("Unable to get multiaddr from IP %s: %v", ip.String(), err)
				return nil
			}

			maddrs = append(maddrs, ma)
		}
	}

	return maddrs
}
