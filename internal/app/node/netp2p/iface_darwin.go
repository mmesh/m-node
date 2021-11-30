//go:build darwin
// +build darwin

package netp2p

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/xlog"
	"x6a.dev/pkg/errors"
)

func (mma *mmAgent) ifUp() error {
	if err := createTUN(); err != nil {
		return errors.Wrapf(err, "[%v] function createTUN()", errors.Trace())
	}

	xlog.Infof("Setting up interface %s", iface.name)

	// set interface parameters
	args := fmt.Sprintf("%s mtu %d -arp up", iface.name, MTU)
	if err := ifconfig(args); err != nil {
		return errors.Wrapf(err, "[%v] function ifconfig()", errors.Trace())
	}

	go mma.readInterface()

	return nil
}

func ifDown() error {
	ifaceName := viper.GetString("agent.iface")

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

	args := fmt.Sprintf("%s down", ifaceName)
	if err := ifconfig(args); err != nil {
		return errors.Wrapf(err, "[%v] function ifconfig()", errors.Trace())
	}

	// args = fmt.Sprintf("%s destroy", ifaceName)
	// if err := ifconfig(args); err != nil {
	// 	return errors.Wrapf(err, "[%v] function ifconfig()", errors.Trace())
	// }

	iface = nil

	return nil
}

func ip4AddrAdd(ipv4 string) error {
	if iface == nil || len(ipv4) == 0 {
		return nil
	}

	if !checkInterfaceAddr(ipv4) {
		return nil
	}

	args := fmt.Sprintf("%s inet %s/32 %s add", iface.name, ipv4, ipv4)

	return ifconfig(args)
}

func ip4AddrDel(ipv4 string) error {
	if iface == nil || len(ipv4) == 0 {
		return nil
	}

	args := fmt.Sprintf("%s inet %s/32 %s delete", iface.name, ipv4, ipv4)

	return ifconfig(args)
}

func ip6AddrAdd(ipv6 string) error {
	if iface == nil || len(ipv6) == 0 {
		return nil
	}

	if !checkInterfaceAddr(ipv6) {
		return nil
	}

	args := fmt.Sprintf("%s inet6 %s/128 add", iface.name, ipv6)

	return ifconfig(args)
}

func ip6AddrDel(ipv6 string) error {
	if iface == nil || len(ipv6) == 0 {
		return nil
	}

	args := fmt.Sprintf("%s inet6 %s/128 delete", iface.name, ipv6)

	return ifconfig(args)
}

func ifconfig(sargs string) error {
	return execCommand("ifconfig", sargs)
}
