//go:build linux
// +build linux

package netp2p

import (
	"net"

	"github.com/vishvananda/netlink"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func (m *routeMap) update() error {
	if iface == nil {
		xlog.Alert("Unable to update interface routes: nil pointer")
		return nil
	}

	ifcLink, err := netlink.LinkByName(iface.name)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	sysRouteList, err := netlink.RouteList(ifcLink, netlink.FAMILY_ALL)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.RouteGet()", errors.Trace())
	}

	for r := range m.routes {
		m.routes[r] = false

		_, dst, err := net.ParseCIDR(r)
		if err != nil {
			return errors.Wrapf(err, "[%v] function net.ParseCIDR()", errors.Trace())
		}

		route := &netlink.Route{
			LinkIndex: ifcLink.Attrs().Index,
			Dst:       dst,
		}

		for _, sysRoute := range sysRouteList {
			if route.Dst.String() == sysRoute.Dst.String() && route.LinkIndex == sysRoute.LinkIndex {
				// route is configured
				m.routes[r] = true
			}
		}
	}

	return nil
}

func routeAdd(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	ifcLink, err := netlink.LinkByName(iface.name)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	_, dst, err := net.ParseCIDR(r)
	if err != nil {
		return errors.Wrapf(err, "[%v] function net.ParseCIDR()", errors.Trace())
	}

	route := &netlink.Route{
		LinkIndex: ifcLink.Attrs().Index,
		Dst:       dst,
	}

	if err := netlink.RouteAdd(route); err != nil {
		return errors.Wrapf(err, "[%v] function netlink.RouteAdd()", errors.Trace())
	}

	xlog.Infof("Added route: %s via %s", r, iface.name)

	return nil
}

func routeDel(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	ifcLink, err := netlink.LinkByName(iface.name)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	_, dst, err := net.ParseCIDR(r)
	if err != nil {
		return errors.Wrapf(err, "[%v] function net.ParseCIDR()", errors.Trace())
	}

	route := &netlink.Route{
		LinkIndex: ifcLink.Attrs().Index,
		Dst:       dst,
	}

	sysRouteList, err := netlink.RouteList(ifcLink, netlink.FAMILY_ALL)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.RouteGet()", errors.Trace())
	}

	for _, sysRoute := range sysRouteList {
		if route.Dst.String() == sysRoute.Dst.String() && route.LinkIndex == sysRoute.LinkIndex {
			if err := netlink.RouteDel(route); err != nil {
				return errors.Wrapf(err, "[%v] function netlink.RouteDel()", errors.Trace())
			}

			xlog.Infof("Deleted route: %s via %s", r, iface.name)
		}
	}

	return nil
}
