//go:build linux
// +build linux

package router

import (
	"net"

	"github.com/vishvananda/netlink"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *router) updateLocalRoutes() error {
	if !r.localForwarding {
		return nil
	}

	if r.networkInterface == nil {
		xlog.Alert("Unable to update interface routes: nil pointer")
		return nil
	}

	ifcLink, err := netlink.LinkByName(r.networkInterface.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	sysRouteList, err := netlink.RouteList(ifcLink, netlink.FAMILY_ALL)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.RouteGet()", errors.Trace())
	}

	for ipDst := range r.routes.local {
		r.routes.local[ipDst] = false

		_, dst, err := net.ParseCIDR(ipDst.string())
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
				r.routes.local[ipDst] = true
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

	ifcLink, err := netlink.LinkByName(ni.devName())
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

	xlog.Infof("Added route: %s via %s", r, ni.devName())

	return nil
}

func (ni *networkInterface) routeDel(ipDst cidrIPDst) error {
	r := ipDst.string()

	if len(r) == 0 {
		return nil
	}

	ifcLink, err := netlink.LinkByName(ni.devName())
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

			xlog.Infof("Deleted route: %s via %s", r, ni.devName())
		}
	}

	return nil
}
