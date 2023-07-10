//go:build linux
// +build linux

package router

import (
	"github.com/songgao/water"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const MTU int = 65535 // TUN interface, so only plain IP packet, no ethernet header

type networkInterface struct {
	ifc     *water.Interface
	name    string
	closeCh chan struct{}
}

func createTUN(ifcName string) (*networkInterface, error) {
	config := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name:       ifcName,
			Persist:    false,
			MultiQueue: true,
		},
	}

	// create TUN interface
	dev, err := water.New(config)
	if err != nil {
		xlog.Errorf("Unable to allocate interface: %v", err)
		return nil, errors.Wrapf(err, "[%v] function water.New()", errors.Trace())
	}

	return &networkInterface{
		ifc:     dev,
		name:    dev.Name(),
		closeCh: make(chan struct{}),
	}, nil
}

func (ni *networkInterface) devName() string {
	if len(ni.name) == 0 {
		ni.name = ni.ifc.Name()
	}

	return ni.name
}

func (ni *networkInterface) close() error {
	return ni.ifc.Close()
}

func (ni *networkInterface) read(buff []byte) (int, error) {
	return ni.ifc.Read(buff)
}

func (ni *networkInterface) write(buff []byte) (int, error) {
	return ni.ifc.Write(buff)
}

/*
func (ni *networkInterface) dev() *water.Interface {
	return ni.ifc
}
*/
