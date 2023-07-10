//go:build windows
// +build windows

package router

import (
	"golang.zx2c4.com/wireguard/tun"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const MTU int = 65535 // TUN interface, so only plain IP packet, no ethernet header

const tunOffset int = 0 // extra bytes for packet header (IFF_NO_PI)

type networkInterface struct {
	ifc     tun.Device
	name    string
	closeCh chan struct{}
}

func createTUN(ifcName string) (*networkInterface, error) {
	dev, err := tun.CreateTUN(ifcName, MTU)
	if err != nil {
		xlog.Errorf("Unable to create interface: %v", err)
		return nil, errors.Wrapf(err, "[%v] function tun.CreateTUN()", errors.Trace())
	}

	devName, err := dev.Name()
	if err != nil {
		xlog.Errorf("Unable to get interface name: %v", err)
		return nil, errors.Wrapf(err, "[%v] function dev.Name()", errors.Trace())
	}

	xlog.Infof("Configured interface %s (tunnel %s)", ifcName, devName)

	return &networkInterface{
		ifc:     dev,
		name:    devName,
		closeCh: make(chan struct{}),
	}, nil
}

func (ni *networkInterface) devName() string {
	return ni.name
}

func (ni *networkInterface) close() error {
	return ni.ifc.Close()
}

func (ni *networkInterface) read(buff []byte) (int, error) {
	return ni.ifc.Read(buff, tunOffset)
}

func (ni *networkInterface) write(buff []byte) (int, error) {
	return ni.ifc.Write(buff, tunOffset)
}

/*
func (ni *networkInterface) dev() tun.Device {
	return ni.ifc
}
*/
