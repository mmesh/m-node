//go:build windows
// +build windows

package netp2p

import (
	"fmt"
	"strings"

	"x6a.dev/pkg/xlog"
)

func (m *routeMap) update() error {
	if iface == nil {
		xlog.Alert("Unable to update interface routes: nil pointer")
		return nil
	}

	// TODO

	return nil
}

func routeAdd(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	var args string

	if strings.Contains(r, ":") {
		args = fmt.Sprintf("interface ipv6 add route prefix=%s interface=\"%s\" store=active", r, iface.name)
	} else {
		args = fmt.Sprintf("interface ipv4 add route prefix=%s interface=\"%s\" store=active", r, iface.name)
	}

	if err := netsh(args); err != nil {
		// xlog.Warn(err)
		// return errors.Wrapf(err, "[%v] function netsh()", errors.Trace())
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
		args = fmt.Sprintf("interface ipv6 delete route prefix=%s interface=\"%s\" store=active", r, iface.name)
	} else {
		args = fmt.Sprintf("interface ipv4 delete route prefix=%s interface=\"%s\" store=active", r, iface.name)
	}

	if err := netsh(args); err != nil {
		// xlog.Warn(err)
		// return errors.Wrapf(err, "[%v] function netsh()", errors.Trace())
		return nil
	}

	xlog.Infof("Deleted route: %s via %s", r, iface.name)

	return nil
}
