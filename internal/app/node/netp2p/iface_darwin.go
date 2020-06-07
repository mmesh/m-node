// +build darwin

package netp2p

import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"github.com/lorenzosaino/go-sysctl"
	"github.com/songgao/water"
	"github.com/spf13/viper"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func (mma *mmAgent) ifUp() error {
	if iface != nil {
		return nil
	}

	ifaceName := viper.GetString("agent.iface")

	config := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: ifaceName,
			// Persist:    false,
			// MultiQueue: true,
		},
	}

	// create TUN interface
	ifc, err := water.New(config)
	if err != nil {
		xlog.Alertf("Unable to allocate interface: %v", err)
		return errors.Wrapf(err, "[%v] function water.New()", errors.Trace())
	}

	iface = ifc

	// set interface parameters
	if err := execIPcmd("link", "set", "dev", iface.Name(), "multicast", "off"); err != nil {
		return errors.Wrapf(err, "[%v] function execIPcmd()", errors.Trace())
	}
	if err := execIPcmd("link", "set", "dev", iface.Name(), "allmulticast", "off"); err != nil {
		return errors.Wrapf(err, "[%v] function execIPcmd()", errors.Trace())
	}
	mtu := fmt.Sprintf("%d", MTU)
	if err := execIPcmd("link", "set", "dev", iface.Name(), "mtu", mtu); err != nil {
		return errors.Wrapf(err, "[%v] function execIPcmd()", errors.Trace())
	}
	if err := execIPcmd("link", "set", "dev", iface.Name(), "up"); err != nil {
		return errors.Wrapf(err, "[%v] function execIPcmd()", errors.Trace())
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

	addrs, err := ifc.Addrs()
	if err != nil {
		xlog.Errorf("Unable to get interface %s addrs: %v", ifaceName, err)
		return nil
	}

	if len(addrs) == 0 {
		xlog.Infof("No address configured on network interface %s", ifaceName)
		return nil
	}

	if err := execIPcmd("link", "set", "dev", ifaceName, "down"); err != nil {
		return errors.Wrapf(err, "[%v] function execIPcmd()", errors.Trace())
	}

	if err := execIPcmd("link", "delete", "dev", ifaceName); err != nil {
		return errors.Wrapf(err, "[%v] function execIPcmd()", errors.Trace())
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

	return execIPcmd("addr", "add", ipv4+"/32", "dev", iface.Name())
}

func ip4AddrDel(ipv4 string) error {
	if iface == nil || len(ipv4) == 0 {
		return nil
	}

	return execIPcmd("addr", "del", ipv4+"/32", "dev", iface.Name())
}

func ip6AddrAdd(ipv6 string) error {
	if iface == nil || len(ipv6) == 0 {
		return nil
	}

	if !checkInterfaceAddr(ipv6) {
		return nil
	}

	v, err := sysctl.Get("net.ipv6.conf." + iface.Name() + ".disable_ipv6")
	if err != nil {
		xlog.Errorf("Unable to get sysctl ipv6 config: %v", err)
		return errors.Wrapf(err, "[%v] function sysctl.Get()", errors.Trace())
	}

	if v == "1" {
		if err := sysctl.Set("net.ipv6.conf."+iface.Name()+".disable_ipv6", "0"); err != nil {
			xlog.Errorf("Unable to enable ipv6 via sysctl: %v", err)
			return errors.Wrapf(err, "[%v] function sysctl.Set()", errors.Trace())
		}
	}

	return execIPcmd("-6", "addr", "add", ipv6+"/128", "dev", iface.Name())
}

func ip6AddrDel(ipv6 string) error {
	if iface == nil || len(ipv6) == 0 {
		return nil
	}

	return execIPcmd("-6", "addr", "del", ipv6+"/128", "dev", iface.Name())
}

func routeAdd(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	if err := exec.Command("route", "add", "-net", r, "-interface", iface.Name()).Run(); err != nil {
		return errors.Wrapf(err, "[%v] function exec.Command()", errors.Trace())
	}

	xlog.Infof("Added route: %s via %s", r, iface.Name())

	return nil
}

func routeDel(r string) error {
	if iface == nil || len(r) == 0 {
		return nil
	}

	if err := exec.Command("route", "delete", "-net", r, "-interface", iface.Name()).Run(); err != nil {
		return errors.Wrapf(err, "[%v] function exec.Command()", errors.Trace())
	}

	xlog.Infof("Deleted route: %s via %s", r, iface.Name())

	return nil
}

func execIPcmd(args ...string) error {
	cmd := exec.Command("/usr/local/bin/ip", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
