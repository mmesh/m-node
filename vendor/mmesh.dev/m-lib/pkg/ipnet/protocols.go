package ipnet

import (
	"fmt"
	"strings"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

const (
	NetworkProtocolNotSupported string = "-"
	NetworkProtocolIPv4TCP      string = "tcp"
	NetworkProtocolIPv6TCP      string = "tcp6"
	NetworkProtocolIPv4UDP      string = "udp"
	NetworkProtocolIPv6UDP      string = "udp6"
	// NetworkProtocolANY          string = "any"
)

func IPProtocol(proto string) layers.IPProtocol {
	switch strings.ToUpper(proto) {
	case topology.Protocol_TCP.String():
		return layers.IPProtocolTCP
	case topology.Protocol_UDP.String():
		return layers.IPProtocolUDP
	case strings.ToUpper(topology.Protocol_ICMPv4.String()):
		return layers.IPProtocolICMPv4
	case strings.ToUpper(topology.Protocol_ICMPv6.String()):
		return layers.IPProtocolICMPv6
	}

	return layers.IPProtocolIPv4
}

func validIPProtocol(proto topology.Protocol) bool {
	if proto == topology.Protocol_ANY {
		return true
	}

	if IPProtocol(proto.String()) != layers.IPProtocolIPv4 {
		return true
	}

	return false
}

func GetNetworkProtocol(af AddressFamily, proto layers.IPProtocol) (string, error) {
	switch proto {
	case layers.IPProtocolTCP:
		switch af {
		case AddressFamilyIPv4:
			return NetworkProtocolIPv4TCP, nil
		case AddressFamilyIPv6:
			return NetworkProtocolIPv6TCP, nil
		}
	case layers.IPProtocolUDP:
		switch af {
		case AddressFamilyIPv4:
			return NetworkProtocolIPv4UDP, nil
		case AddressFamilyIPv6:
			return NetworkProtocolIPv6UDP, nil
		}
	}

	return NetworkProtocolNotSupported, fmt.Errorf("protocol %v not supported", proto)
}

func GetPortName(proto layers.IPProtocol, p uint16) string {
	switch proto {
	case layers.IPProtocolTCP:
		return layers.TCPPort(p).String()
	case layers.IPProtocolUDP:
		return layers.UDPPort(p).String()
	}

	return fmt.Sprintf("port-%d", p)
}
