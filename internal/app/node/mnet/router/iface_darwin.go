//go:build darwin
// +build darwin

package router

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/utils"
)

func (r *router) ifUp() error {
	ifcName := viper.GetString("iface")

	if !strings.HasPrefix(ifcName, "utun") {
		ifcName = "utun7"
	}

	ni, err := createTUN(ifcName)
	if err != nil {
		return errors.Wrapf(err, "[%v] function createTUN()", errors.Trace())
	}

	xlog.Infof("Setting up interface %s", ni.devName())

	// set interface parameters
	args := fmt.Sprintf("%s mtu %d -arp up", ni.devName(), MTU)
	if err := utils.Ifconfig(args); err != nil {
		return errors.Wrapf(err, "[%v] function utils.Ifconfig()", errors.Trace())
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

	args := fmt.Sprintf("%s down", ifaceName)
	if err := utils.Ifconfig(args); err != nil {
		return errors.Wrapf(err, "[%v] function utils.Ifconfig()", errors.Trace())
	}

	// args = fmt.Sprintf("%s destroy", ifaceName)
	// if err := ifconfig(args); err != nil {
	// 	return errors.Wrapf(err, "[%v] function ifconfig()", errors.Trace())
	// }

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

	args := fmt.Sprintf("%s inet %s/32 %s add", ni.devName(), ipv4, ipv4)

	return utils.Ifconfig(args)
}

func (ni *networkInterface) ip4AddrDel(ipv4 string) error {
	if len(ipv4) == 0 {
		return nil
	}

	args := fmt.Sprintf("%s inet %s/32 %s delete", ni.devName(), ipv4, ipv4)

	return utils.Ifconfig(args)
}

func (ni *networkInterface) ip6AddrAdd(ipv6 string) error {
	if len(ipv6) == 0 {
		return nil
	}

	if !ni.checkInterfaceAddr(ipv6) {
		return nil
	}

	args := fmt.Sprintf("%s inet6 %s/128 add", ni.devName(), ipv6)

	return utils.Ifconfig(args)
}

func (ni *networkInterface) ip6AddrDel(ipv6 string) error {
	if len(ipv6) == 0 {
		return nil
	}

	args := fmt.Sprintf("%s inet6 %s/128 delete", ni.devName(), ipv6)

	return utils.Ifconfig(args)
}
