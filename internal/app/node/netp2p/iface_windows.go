//go:build windows
// +build windows

package netp2p

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func (mma *mmAgent) ifUp() error {
	if err := createTUN(); err != nil {
		return errors.Wrapf(err, "[%v] function createTUN()", errors.Trace())
	}

	xlog.Infof("Setting up interface %s", iface.name)

	// set interface parameters
	args := fmt.Sprintf("interface ipv4 set subinterface \"%s\" mtu=%d store=active", iface.name, MTU)
	if err := netsh(args); err != nil {
		return errors.Wrapf(err, "[%v] function netsh()", errors.Trace())
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

	args := fmt.Sprintf("interface ipv4 set address name=\"%s\" source=static addr=%s/32 mask=255.255.255.255 gateway=none store=active", iface.name, ipv4)

	return netsh(args)
}

func ip4AddrDel(ipv4 string) error {
	if iface == nil || len(ipv4) == 0 {
		return nil
	}

	args := fmt.Sprintf("interface ipv4 delete address name=\"%s\" addr=%s gateway=all store=active", iface.name, ipv4)

	return netsh(args)
}

func ip6AddrAdd(ipv6 string) error {
	if iface == nil || len(ipv6) == 0 {
		return nil
	}

	if !checkInterfaceAddr(ipv6) {
		return nil
	}

	args := fmt.Sprintf("interface ipv6 set address interface=\"%s\" type=unicast address=%s/128 store=active", iface.name, ipv6)

	return netsh(args)
}

func ip6AddrDel(ipv6 string) error {
	if iface == nil || len(ipv6) == 0 {
		return nil
	}

	args := fmt.Sprintf("interface ipv6 delete address interface=\"%s\" address=%s store=active", iface.name, ipv6)

	return netsh(args)
}

func netsh(sargs string) error {
	return execCommand("netsh", sargs)
}
