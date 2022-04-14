//go:build darwin
// +build darwin

package netp2p

import (
	"fmt"
	"net"
	"strings"
	"syscall"

	"golang.org/x/net/route"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (m *routeMap) update() error {
	if iface == nil {
		xlog.Alert("Unable to update interface routes: nil pointer")
		return nil
	}

	ifc, err := net.InterfaceByName(iface.name)
	if err != nil {
		return errors.Wrapf(err, "[%v] function net.InterfaceByName(()", errors.Trace())
	}

	rib, err := route.FetchRIB(syscall.AF_UNSPEC, route.RIBTypeRoute, 0)
	if err != nil {
		return errors.Wrapf(err, "[%v] function route.FetchRIB()", errors.Trace())
	}

	msgs, err := route.ParseRIB(route.RIBTypeRoute, rib)
	if err != nil {
		return errors.Wrapf(err, "[%v] function route.ParseRIB()", errors.Trace())
	}

	for r := range m.routes {
		m.routes[r] = false

		ipAddr, _, err := net.ParseCIDR(r)
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
				switch a := msg.Addrs[syscall.RTAX_DST].(type) {
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

func routeAdd(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	var args string

	if strings.Contains(r, ":") {
		prefix := strings.Split(r, "/")[0]
		prefixlen := strings.Split(r, "/")[1]
		args = fmt.Sprintf("-n add -inet6 -net %s -prefixlen %s -interface %s", prefix, prefixlen, iface.name)
	} else {
		args = fmt.Sprintf("-n add -inet -net %s -interface %s", r, iface.name)
	}

	if err := routeCmd(args); err != nil {
		xlog.Warnf("Unable to add route: %v", err)
		return nil
	}

	xlog.Infof("Added route: %s via %s", r, iface.name)

	return nil
}

func routeDel(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	var args string

	if strings.Contains(r, ":") {
		args = fmt.Sprintf("-n delete -inet6 -net %s -interface %s", r, iface.name)
	} else {
		args = fmt.Sprintf("-n delete -inet -net %s -interface %s", r, iface.name)
	}

	if err := routeCmd(args); err != nil {
		xlog.Warnf("Unable to remove route: %v", err)
		return nil
	}

	xlog.Infof("Deleted route: %s via %s", r, iface.name)

	return nil
}

func routeCmd(sargs string) error {
	return execCommand("route", sargs)
}
