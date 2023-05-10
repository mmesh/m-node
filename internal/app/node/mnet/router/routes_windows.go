//go:build windows
// +build windows

package router

import (
	"fmt"
	"strings"

	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/utils"
)

func (ni *networkInterface) routeAdd(ipDst cidrIPDst) error {
	r := ipDst.string()

	if len(r) == 0 {
		return nil
	}

	var args string

	if strings.Contains(r, ":") {
		args = fmt.Sprintf("interface ipv6 add route prefix=%s interface=\"%s\" store=active", r, ni.devName())
	} else {
		args = fmt.Sprintf("interface ipv4 add route prefix=%s interface=\"%s\" store=active", r, ni.devName())
	}

	if err := utils.Netsh(args); err != nil {
		// xlog.Warn(err)
		// return errors.Wrapf(err, "[%v] function utils.Netsh()", errors.Trace())
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
		args = fmt.Sprintf("interface ipv6 delete route prefix=%s interface=\"%s\" store=active", r, ni.devName())
	} else {
		args = fmt.Sprintf("interface ipv4 delete route prefix=%s interface=\"%s\" store=active", r, ni.devName())
	}

	if err := utils.Netsh(args); err != nil {
		// xlog.Warn(err)
		// return errors.Wrapf(err, "[%v] function utils.Netsh()", errors.Trace())
		return nil
	}

	xlog.Infof("Deleted route: %s via %s", r, ni.devName())

	return nil
}

/*
func (r *router) updateLocalRoutes() error {
	if !r.localForwarding {
		return nil
	}

	if r.networkInterface == nil {
		xlog.Alert("Unable to update interface routes: nil pointer")
		return nil
	}

	// TODO

	return nil
}
*/
