package ipnet

import (
	"fmt"
	"strings"

	"github.com/google/gopacket/layers"
)

const (
	NetworkProtocolNotSupported string = "-"
	NetworkProtocolIPv4TCP      string = "tcp"
	NetworkProtocolIPv6TCP      string = "tcp6"
	NetworkProtocolIPv4UDP      string = "udp"
	NetworkProtocolIPv6UDP      string = "udp6"
	NetworkProtocolANY          string = "any"
)

func IPProtocol(proto string) layers.IPProtocol {
	if strings.Contains(strings.ToLower(proto), "tcp") {
		return layers.IPProtocolTCP
	}

	if strings.Contains(strings.ToLower(proto), "udp") {
		return layers.IPProtocolUDP
	}

	if strings.Contains(strings.ToLower(proto), "gre") {
		return layers.IPProtocolGRE
	}

	if strings.Contains(strings.ToLower(proto), "esp") {
		return layers.IPProtocolESP
	}

	if strings.Contains(strings.ToLower(proto), "ah") {
		return layers.IPProtocolAH
	}

	if strings.Contains(strings.ToLower(proto), "ospf") {
		return layers.IPProtocolOSPF
	}

	if strings.Contains(strings.ToLower(proto), "ipip") {
		return layers.IPProtocolIPIP
	}

	if strings.Contains(strings.ToLower(proto), "vrrp") {
		return layers.IPProtocolVRRP
	}

	if strings.Contains(strings.ToLower(proto), "igmp") {
		return layers.IPProtocolIGMP
	}

	if strings.Contains(strings.ToLower(proto), "icmpv4") {
		return layers.IPProtocolICMPv4
	}

	if strings.Contains(strings.ToLower(proto), "icmpv6") {
		return layers.IPProtocolICMPv6
	}

	return layers.IPProtocolIPv4
}

func validIPProtocol(proto string) bool {
	if strings.ToLower(proto) == NetworkProtocolANY {
		return true
	}

	if IPProtocol(proto) != layers.IPProtocolIPv4 {
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
