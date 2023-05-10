//go:build linux
// +build linux

package router

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func routeExists(ipDst cidrIPDst, ifName string) (bool, error) {
	ifcLink, err := netlink.LinkByName(ifName)
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	sysRouteList, err := netlink.RouteList(ifcLink, netlink.FAMILY_ALL)
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function netlink.RouteGet()", errors.Trace())
	}

	_, dst, err := net.ParseCIDR(ipDst.string())
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function net.ParseCIDR()", errors.Trace())
	}

	route := &netlink.Route{
		LinkIndex: ifcLink.Attrs().Index,
		Dst:       dst,
	}

	for _, sysRoute := range sysRouteList {
		if route.Dst.String() == sysRoute.Dst.String() && route.LinkIndex == sysRoute.LinkIndex {
			// route is configured
			return true, nil
		}
	}

	return false, nil
}

func (ni *networkInterface) routeAdd(ipDst cidrIPDst) error {
	ok, err := routeExists(ipDst, ni.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function routeExists()", errors.Trace())
	}
	if ok {
		return fmt.Errorf("route to %s already exists", ipDst.string())
	}

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

/*
func (r *router) routeExists(ipDst cidrIPDst) (bool, error) {
	if !r.localForwarding {
		return false, fmt.Errorf("localForwarding disabled")
	}

	if r.networkInterface == nil {
		return false, fmt.Errorf("invalid networkInterface")
	}

	ifcLink, err := netlink.LinkByName(r.networkInterface.devName())
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	sysRouteList, err := netlink.RouteList(ifcLink, netlink.FAMILY_ALL)
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function netlink.RouteGet()", errors.Trace())
	}

	_, dst, err := net.ParseCIDR(ipDst.string())
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function net.ParseCIDR()", errors.Trace())
	}

	route := &netlink.Route{
		LinkIndex: ifcLink.Attrs().Index,
		Dst:       dst,
	}

	for _, sysRoute := range sysRouteList {
		if route.Dst.String() == sysRoute.Dst.String() && route.LinkIndex == sysRoute.LinkIndex {
			// route is configured
			return true, nil
		}
	}

	return false, nil
}

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
*/
