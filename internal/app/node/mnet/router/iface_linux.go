//go:build linux
// +build linux

package router

import (
	"net"

	"github.com/lorenzosaino/go-sysctl"
	"github.com/spf13/viper"
	"github.com/vishvananda/netlink"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *router) ifUp() error {
	ifcName := viper.GetString("agent.iface")

	ni, err := createTUN(ifcName)
	if err != nil {
		return errors.Wrapf(err, "[%v] function createTUN()", errors.Trace())
	}

	xlog.Infof("Setting up interface %s", ni.devName())

	// set interface parameters
	ifcLink, err := netlink.LinkByName(ni.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	if err := netlink.LinkSetAllmulticastOff(ifcLink); err != nil {
		xlog.Alertf("Unable to set interface allmulticast off: %v", err)
		// return nil, errors.Wrapf(err, "[%v] function netlink.LinkSetAllmulticastOff()", errors.Trace())
	}

	if err := netlink.LinkSetMTU(ifcLink, MTU); err != nil {
		xlog.Alertf("Unable to set interface MTU: %v", err)
		// return nil, errors.Wrapf(err, "[%v] function netlink.LinkSetMTU()", errors.Trace())
	}

	if err := netlink.LinkSetUp(ifcLink); err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkSetUp()", errors.Trace())
	}

	r.networkInterface = ni

	go r.readInterface()

	return nil
}

func (r *router) ifDown() error {
	var ifaceName string

	if r.networkInterface == nil {
		ifaceName = viper.GetString("agent.iface")
	} else {
		ifaceName = r.networkInterface.devName()
	}

	ifc, err := net.InterfaceByName(ifaceName)
	if err != nil {
		xlog.Infof("Network interface %s not configured", ifaceName)
		return nil
	}

	xlog.Infof("Bringing down interface %s", ifaceName)

	addrs, err := ifc.Addrs()
	if err != nil {
		xlog.Errorf("Unable to get interface %s addrs: %v", ifaceName, err)
		return nil
	}

	if len(addrs) == 0 {
		xlog.Infof("No address configured on network interface %s", ifaceName)
		return nil
	}

	ifcLink, err := netlink.LinkByName(ifaceName)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	if err := netlink.LinkSetDown(ifcLink); err != nil {
		xlog.Alertf("Unable to set interface down: %v", err)
		// return errors.Wrapf(err, "[%v] function netlink.LinkSetDown()", errors.Trace())
	}

	if err := netlink.LinkDel(ifcLink); err != nil {
		xlog.Alertf("Unable to delete interface: %v", err)
		return errors.Wrapf(err, "[%v] function netlink.LinkDel()", errors.Trace())
	}

	if r.networkInterface != nil {
		return r.networkInterface.close()
	}

	return nil
}

func (ni *networkInterface) ip4AddrAdd(ipv4 string) error {
	if len(ipv4) == 0 {
		return nil
	}

	if !ni.checkInterfaceAddr(ipv4) {
		return nil
	}

	ifcLink, err := netlink.LinkByName(ni.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	ipAddr, err := netlink.ParseAddr(ipv4 + "/32")
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.ParseAddr()", errors.Trace())
	}

	if err := netlink.AddrAdd(ifcLink, ipAddr); err != nil {
		return errors.Wrapf(err, "[%v] function netlink.AddrAdd()", errors.Trace())
	}

	return nil
}

func (ni *networkInterface) ip4AddrDel(ipv4 string) error {
	if len(ipv4) == 0 {
		return nil
	}

	ifcLink, err := netlink.LinkByName(ni.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	addr, err := netlink.ParseAddr(ipv4 + "/32")
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.ParseAddr()", errors.Trace())
	}

	if err := netlink.AddrDel(ifcLink, addr); err != nil {
		return errors.Wrapf(err, "[%v] function netlink.AddrDel()", errors.Trace())
	}

	return nil
}

func (ni *networkInterface) ip6AddrAdd(ipv6 string) error {
	if len(ipv6) == 0 {
		return nil
	}

	if !ni.checkInterfaceAddr(ipv6) {
		return nil
	}

	ifaceName := ni.devName()

	v, err := sysctl.Get("net.ipv6.conf." + ifaceName + ".disable_ipv6")
	if err != nil {
		xlog.Errorf("Unable to get sysctl ipv6 config: %v", err)
		return errors.Wrapf(err, "[%v] function sysctl.Get()", errors.Trace())
	}

	if v == "1" {
		if err := sysctl.Set("net.ipv6.conf."+ifaceName+".disable_ipv6", "0"); err != nil {
			xlog.Errorf("Unable to enable ipv6 via sysctl: %v", err)
			return errors.Wrapf(err, "[%v] function sysctl.Set()", errors.Trace())
		}
	}

	ifcLink, err := netlink.LinkByName(ifaceName)
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	addr, err := netlink.ParseAddr(ipv6 + "/128")
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.ParseAddr()", errors.Trace())
	}

	if err := netlink.AddrAdd(ifcLink, addr); err != nil {
		return errors.Wrapf(err, "[%v] function netlink.AddrAdd()", errors.Trace())
	}

	return nil
}

func (ni *networkInterface) ip6AddrDel(ipv6 string) error {
	if len(ipv6) == 0 {
		return nil
	}

	ifcLink, err := netlink.LinkByName(ni.devName())
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.LinkByName()", errors.Trace())
	}

	addr, err := netlink.ParseAddr(ipv6 + "/128")
	if err != nil {
		return errors.Wrapf(err, "[%v] function netlink.ParseAddr()", errors.Trace())
	}

	if err := netlink.AddrDel(ifcLink, addr); err != nil {
		return errors.Wrapf(err, "[%v] function netlink.AddrDel()", errors.Trace())
	}

	return nil
}
