// +build darwin

package netp2p

import (
	"net"
	"syscall"

	"golang.org/x/net/route"
	"x6a.dev/pkg/errors"
)

func (m *routeMap) update() error {
	if iface == nil {
		return nil
	}

	ifc, err := net.InterfaceByName(iface.Name())
	if err != nil {
		return errors.Wrapf(err, "[%v] function net.InterfaceByName(()", errors.Trace())
	}

	// not-tested

	rib, err = route.FetchRIB(syscall.AF_UNSPEC, route.RIBTypeRoute, 0)
	if err != nil {
		return errors.Wrapf(err, "[%v] function route.FetchRIB()", errors.Trace())
	}

	msgs, err := route.ParseRIB(route.RIBTypeRoute, rib)
	if err != nil {
		return errors.Wrapf(err, "[%v] function route.ParseRIB()", errors.Trace())
	}

	for r := range m.routes {
		m.routes[r] = false

		ipAddr, _ := net.ParseCIDR(r)
		if err != nil {
			return errors.Wrapf(err, "[%v] function net.ParseCIDR()", errors.Trace())
		}

		for _, msg := range msgs {
			switch msg := msg.(type) {
			case *route.RouteMessage:
				if ifc.Index != msg.Index {
					continue
				}

				var ip net.IP
				switch a := msg.Addrs[route.RTAX_DST].(type) {
				case *route.Inet4Addr:
					ip = net.IPv4(a.IP[0], a.IP[1], a.IP[2], a.IP[3])
				case *route.Inet6Addr:
					ip = make(net.IP, net.IPv6len)
					copy(ip, a.IP[:])
				}

				if ipAddr.Equal(ip) {
					m.routes[r] = true
				}
			}
		}
	}

	return nil
}
