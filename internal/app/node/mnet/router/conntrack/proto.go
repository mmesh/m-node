package conntrack

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"mmesh.dev/m-lib/pkg/ipnet"
)

// Filter specific proto packet types
func (conn *Connection) ProtoFilter() bool {
	if conn.Proto == layers.IPProtocolICMPv4 || conn.Proto == layers.IPProtocolICMPv6 {
		if conn.invalidICMPTypeRequest() {
			return true // only icmp echo request is permitted, drop the pkt
		}
	}

	return false // proto packet type is valid, accept the pkt
}

func (conn *Connection) ParseProtocol(pkt []byte) error {
	var p gopacket.Packet

	switch conn.AF {
	case ipnet.AddressFamilyIPv4:
		p = gopacket.NewPacket(pkt, layers.LayerTypeIPv4, gopacket.Default)
	case ipnet.AddressFamilyIPv6:
		p = gopacket.NewPacket(pkt, layers.LayerTypeIPv6, gopacket.Default)
	default:
		return fmt.Errorf("unknown IP address family")
	}

	// Get the TCP layer from this packet
	if tcpLayer := p.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		conn.Proto = layers.IPProtocolTCP
		conn.SrcPort = uint16(tcp.SrcPort)
		conn.DstPort = uint16(tcp.DstPort)
		return nil
	}

	// Get the UDP layer from this packet
	if udpLayer := p.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		conn.Proto = layers.IPProtocolUDP
		conn.SrcPort = uint16(udp.SrcPort)
		conn.DstPort = uint16(udp.DstPort)
		return nil
	}

	// Get the ICMPv4 layer from this packet
	if icmp4Layer := p.Layer(layers.LayerTypeICMPv4); icmp4Layer != nil {
		icmp4, _ := icmp4Layer.(*layers.ICMPv4)
		conn.Proto = layers.IPProtocolICMPv4
		conn.ICMPv4TypeCode = icmp4.TypeCode
		return nil
	}

	// Get the ICMPv6 layer from this packet
	if icmp6Layer := p.Layer(layers.LayerTypeICMPv6); icmp6Layer != nil {
		icmp6, _ := icmp6Layer.(*layers.ICMPv6)
		conn.Proto = layers.IPProtocolICMPv6
		conn.ICMPv6TypeCode = icmp6.TypeCode
		return nil
	}

	// Get the GRE layer from this packet
	if greLayer := p.Layer(layers.LayerTypeGRE); greLayer != nil {
		gre, _ := greLayer.(*layers.GRE)
		conn.Proto = layers.IPProtocolGRE
		conn.GREKey = gre.Key
		return nil
	}

	return nil
}
