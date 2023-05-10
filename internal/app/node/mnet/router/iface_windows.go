//go:build windows
// +build windows

package router

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/utils"
)

func (r *router) ifUp() error {
	ifcName := viper.GetString("iface")

	ni, err := createTUN(ifcName)
	if err != nil {
		return errors.Wrapf(err, "[%v] function createTUN()", errors.Trace())
	}

	xlog.Infof("Setting up interface %s", ni.devName())

	// set interface parameters
	args := fmt.Sprintf("interface ipv4 set subinterface \"%s\" mtu=%d store=active", ni.devName(), MTU)
	if err := utils.Netsh(args); err != nil {
		return errors.Wrapf(err, "[%v] function utils.Netsh()", errors.Trace())
	}

	r.networkInterface = ni

	go r.readInterface()

	return nil
}

func (r *router) ifDown() error {
	var ifaceName string

	if r.networkInterface == nil {
		ifaceName = viper.GetString("iface")
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

	args := fmt.Sprintf("interface ipv4 set address name=\"%s\" source=static addr=%s/32 mask=255.255.255.255 gateway=none store=active", ni.devName(), ipv4)

	return utils.Netsh(args)
}

func (ni *networkInterface) ip4AddrDel(ipv4 string) error {
	if len(ipv4) == 0 {
		return nil
	}

	args := fmt.Sprintf("interface ipv4 delete address name=\"%s\" addr=%s gateway=all store=active", ni.devName(), ipv4)

	return utils.Netsh(args)
}

func (ni *networkInterface) ip6AddrAdd(ipv6 string) error {
	if len(ipv6) == 0 {
		return nil
	}

	if !ni.checkInterfaceAddr(ipv6) {
		return nil
	}

	args := fmt.Sprintf("interface ipv6 set address interface=\"%s\" type=unicast address=%s/128 store=active", ni.devName(), ipv6)

	return utils.Netsh(args)
}

func (ni *networkInterface) ip6AddrDel(ipv6 string) error {
	if len(ipv6) == 0 {
		return nil
	}

	args := fmt.Sprintf("interface ipv6 delete address interface=\"%s\" address=%s store=active", ni.devName(), ipv6)

	return utils.Netsh(args)
}
