//go:build darwin
// +build darwin

package router

import (
	"fmt"
	"net"
	"strings"
	"syscall"

	"golang.org/x/net/route"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/utils"
)

func (r *router) updateLocalRoutes() error {
	if !r.localForwarding {
		return nil
	}

	if r.networkInterface == nil {
		xlog.Alert("Unable to update interface routes: nil pointer")
		return nil
	}

	ifc, err := net.InterfaceByName(r.networkInterface.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function net.InterfaceByName(()", errors.Trace())
	}

	localRIB, err := route.FetchRIB(syscall.AF_UNSPEC, route.RIBTypeRoute, 0)
	if err != nil {
		return errors.Wrapf(err, "[%v] function route.FetchRIB()", errors.Trace())
	}

	msgs, err := route.ParseRIB(route.RIBTypeRoute, localRIB)
	if err != nil {
		return errors.Wrapf(err, "[%v] function route.ParseRIB()", errors.Trace())
	}

	for ipDst := range r.routes.local {
		r.routes.local[ipDst] = false

		ipAddr, _, err := net.ParseCIDR(ipDst.string())
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
					r.routes.local[ipDst] = true
				}
			}
		}
	}

	return nil
}

func (ni *networkInterface) routeAdd(ipDst cidrIPDst) error {
	r := ipDst.string()

	if len(r) == 0 {
		return nil
	}

	var args string

	if strings.Contains(r, ":") {
		prefix := strings.Split(r, "/")[0]
		prefixlen := strings.Split(r, "/")[1]
		args = fmt.Sprintf("-n add -inet6 -net %s -prefixlen %s -interface %s", prefix, prefixlen, ni.devName())
	} else {
		args = fmt.Sprintf("-n add -inet -net %s -interface %s", r, ni.devName())
	}

	if err := utils.Route(args); err != nil {
		xlog.Warnf("Unable to add route: %v", err)
		return nil
	}

	xlog.Infof("Added route: %s via %s", r, ni.devName())

	return nil
}

func (ni *networkInterface) routeDel(ipDst cidrIPDst) error {
	r := ipDst.string()

	if len(r) == 0 {
		return nil
	}

	var args string

	if strings.Contains(r, ":") {
		args = fmt.Sprintf("-n delete -inet6 -net %s -interface %s", r, ni.devName())
	} else {
		args = fmt.Sprintf("-n delete -inet -net %s -interface %s", r, ni.devName())
	}

	if err := utils.Route(args); err != nil {
		xlog.Warnf("Unable to remove route: %v", err)
		return nil
	}

	xlog.Infof("Deleted route: %s via %s", r, ni.devName())

	return nil
}
