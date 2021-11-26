//go:build darwin
// +build darwin

package netp2p

import (
	"github.com/songgao/water"
	"github.com/spf13/viper"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

const MTU int = 1400 // TUN interface, so only plain IP packet, no ethernet header + mtu is set to 1300

type mmeshNetworkInterface struct {
	ifc  *water.Interface
	name string
}

var iface *mmeshNetworkInterface

func createTUN() error {
	if iface != nil {
		return nil
	}

	ifcName := viper.GetString("agent.iface")

	config := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: ifcName,
		},
	}

	// create TUN interface
	dev, err := water.New(config)
	if err != nil {
		xlog.Alertf("Unable to allocate interface: %v", err)
		return errors.Wrapf(err, "[%v] function water.New()", errors.Trace())
	}

	devName := dev.Name()

	xlog.Infof("Configured interface %s", devName)

	viper.Set("agent.iface", devName)

	iface = &mmeshNetworkInterface{
		ifc:  dev,
		name: devName,
	}

	return nil
}

/*
func (i *mmeshNetworkInterface) dev() *water.Interface {
	return i.ifc
}

func (i *mmeshNetworkInterface) devName() (string, error) {
	if len(i.name) == 0 {
		i.name = i.ifc.Name()
	}

	return i.name, nil
}
*/

func ifaceClose() error {
	return iface.ifc.Close()
}

func ifaceRead(buff []byte) (int, error) {
	return iface.ifc.Read(buff)
}

func ifaceWrite(buff []byte) (int, error) {
	return iface.ifc.Write(buff)
}
