//go:build windows
// +build windows

package netp2p

import (
	"github.com/spf13/viper"
	"golang.zx2c4.com/wireguard/tun"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

const MTU int = 1400 // TUN interface, so only plain IP packet, no ethernet header + mtu is set to 1300

const tunOffset int = 0 // extra bytes for packet header (IFF_NO_PI)

type mmeshNetworkInterface struct {
	ifc  tun.Device
	name string
}

var iface *mmeshNetworkInterface

func createTUN() error {
	if iface != nil {
		return nil
	}

	ifcName := viper.GetString("agent.iface")

	// create TUN interface
	dev, err := tun.CreateTUN(ifcName, MTU)
	if err != nil {
		xlog.Alertf("Unable to create interface: %v", err)
		return errors.Wrapf(err, "[%v] function tun.CreateTUN()", errors.Trace())
	}

	devName, err := dev.Name()
	if err != nil {
		xlog.Alertf("Unable to get interface name: %v", err)
		return errors.Wrapf(err, "[%v] function dev.Name()", errors.Trace())
	}

	xlog.Infof("Configured interface %s (tunnel %s)", ifcName, devName)

	viper.Set("agent.iface", devName)

	iface = &mmeshNetworkInterface{
		ifc:  dev,
		name: devName,
	}

	return nil
}

/*
func (i *mmeshNetworkInterface) dev() tun.Device {
	return i.ifc
}

func (i *mmeshNetworkInterface) devName() (string, error) {
	if len(i.name) == 0 {
		devName, err := i.ifc.Name()
		if err != nil {
			xlog.Alertf("Unable to get interface name: %v", err)
			return "", errors.Wrapf(err, "[%v] function i.ifc.Name()", errors.Trace())
		}
		i.name = devName
	}

	return i.name, nil
}
*/

func ifaceClose() error {
	return iface.ifc.Close()
}

func ifaceRead(buff []byte) (int, error) {
	return iface.ifc.Read(buff, tunOffset)
}

func ifaceWrite(buff []byte) (int, error) {
	return iface.ifc.Write(buff, tunOffset)
}
