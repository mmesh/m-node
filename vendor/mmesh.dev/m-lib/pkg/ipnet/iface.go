package ipnet

import (
	"fmt"
	"net"

	"mmesh.dev/m-lib/pkg/errors"
)

func getInterfaceHwAddr() (net.HardwareAddr, error) {
	ifcs, err := net.Interfaces()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function net.Interfaces()", errors.Trace())
	}

	for _, ifc := range ifcs {
		if len(ifc.HardwareAddr) == 0 {
			continue
		}

		// Skip loopback interfaces
		if ifc.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Skip locally administered addresses
		// if ifc.HardwareAddr[0]&2 == 2 {
		// 	continue
		// }

		if ifc.Flags&net.FlagUp != 0 {
			return ifc.HardwareAddr, nil
		}
	}

	return nil, fmt.Errorf("mac address not found")
}

/*
func getInterfaceHwAddr(ifaceName string) (net.HardwareAddr, error) {
	if len(ifaceName) == 0 {
		return nil, fmt.Errorf("invalid network interface")
	}

	ifc, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function net.InterfaceByName()", errors.Trace())
	}

	return ifc.HardwareAddr, nil

	// if len(ifc.HardwareAddr.String()) == 0 {
	// 	return nil, fmt.Errorf("invalid hardware addr")
	// }

	// hwAddr, err := net.ParseMAC(ifc.HardwareAddr.String())
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "[%v] function net.ParseMAC()", errors.Trace())
	// }

	// return hwAddr, nil
}
*/
